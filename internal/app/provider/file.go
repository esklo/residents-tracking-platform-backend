package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/file"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	fileRepository "github.com/esklo/residents-tracking-platform-backend/internal/repository/file"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	fileService "github.com/esklo/residents-tracking-platform-backend/internal/service/file"
	"github.com/esklo/residents-tracking-platform-backend/internal/storage/s3"
)

func (s *ServiceProvider) FileRepository() repository.FileRepository {
	if s.fileRepository == nil {
		s.fileRepository = fileRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.fileRepository
}

func (s *ServiceProvider) FileService() service.FileService {
	if s.fileService == nil {
		s.fileService = fileService.NewService(
			s.FileRepository(),
			s.FileStorage(),
			s.GetLogger(),
		)
	}

	return s.fileService
}

func (s *ServiceProvider) FileImpl() *file.Implementation {
	if s.fileImpl == nil {
		s.fileImpl = file.NewImplementation(s.FileService(), s.AuthService(), s.GetLogger())
	}

	return s.fileImpl
}

func (s *ServiceProvider) FileStorage() *s3.Storage {
	if s.fileStorage == nil {
		cfg := s.S3Config()
		s.fileStorage = s3.NewStorage(cfg.Endpoint(), cfg.Region(), cfg.KeyId(), cfg.AccessKey(), cfg.Bucket(), s.GetLogger())
	}

	return s.fileStorage
}
