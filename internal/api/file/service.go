package file

import (
	"context"
	"github.com/esklo/residents-tracking-platform/gen/proto/empty"
	proto "github.com/esklo/residents-tracking-platform/gen/proto/file"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/esklo/residents-tracking-platform/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Implementation struct {
	proto.UnimplementedFileServiceServer
	fileService service.FileService
	authService service.AuthService
}

func NewImplementation(fileService service.FileService, authService service.AuthService) *Implementation {
	return &Implementation{
		fileService: fileService,
		authService: authService,
	}
}

func (i Implementation) Upload(ctx context.Context, req *proto.UploadRequest) (*proto.File, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	file, err := i.fileService.Upload(ctx, req.Filename, req.File)
	if err != nil {
		return nil, errors.Wrap(err, "can not upload file via service")
	}
	protoFile, err := file.ToProto()
	if err != nil {
		return nil, errors.Wrap(err, "can not convert file to proto")
	}
	return protoFile, nil
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.File, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	fileId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, model.ErrorInvalidUUID
	}
	file, err := i.fileService.GetById(ctx, &fileId)
	if err != nil {
		return nil, err
	}
	protoFile, err := file.ToProto()
	if err != nil {
		return nil, err
	}
	return protoFile, nil
}

func (i Implementation) Delete(ctx context.Context, req *proto.ByIdRequest) (*empty.Empty, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	fileId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, model.ErrorInvalidUUID
	}
	if err := i.fileService.Delete(ctx, &fileId); err != nil {
		return nil, errors.Wrap(err, "can no delete file via service")
	}
	return &empty.Empty{}, nil
}
