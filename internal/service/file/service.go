package file

import (
	"bytes"
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/esklo/residents-tracking-platform-backend/internal/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var _ def.FileService = (*Service)(nil)

type Service struct {
	fileRepository repository.FileRepository
	fileStorage    storage.Storage
	logger         *zap.Logger
}

func NewService(
	fileRepository repository.FileRepository,
	fileStorage storage.Storage,
	logger *zap.Logger,
) *Service {
	return &Service{
		fileRepository: fileRepository,
		fileStorage:    fileStorage,
		logger:         logger,
	}
}

func (s Service) Upload(ctx context.Context, filename string, data []byte) (*model.File, error) {
	filePath, mime, err := s.fileStorage.PutFile(bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrap(err, "can not put file to storage")
	}

	fileUrl, err := s.fileStorage.GetFileUrl(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "can not retrieve file url while uploading")
	}
	file := &model.File{
		FileName:  filename,
		MimeType:  mime.String(),
		Extension: mime.Extension(),
		Path:      filePath,
		PublicUrl: fileUrl,
	}
	return s.fileRepository.Create(ctx, file)
}

func (s Service) GetById(ctx context.Context, id *uuid.UUID) (*model.File, error) {
	fileModel, err := s.fileRepository.GetByID(ctx, id.String())
	if err != nil {
		return nil, err
	}
	fileUrl, err := s.fileStorage.GetFileUrl(fileModel.Path)
	if err != nil {
		return nil, errors.Wrap(err, "can not retrieve file url")
	}
	fileModel.PublicUrl = fileUrl

	return fileModel, nil
}

func (s Service) Delete(ctx context.Context, id *uuid.UUID) error {
	fileModel, err := s.fileRepository.GetByID(ctx, id.String())
	if err != nil {
		return errors.Wrap(err, "can not get file by id")
	}
	if err := s.fileRepository.Delete(ctx, id.String()); err != nil {
		return errors.Wrap(err, "can not delete file from repository")
	}
	if err := s.fileStorage.DeleteFile(fileModel.Path); err != nil {
		return errors.Wrap(err, "can not delete file from storage")
	}

	return nil
}
