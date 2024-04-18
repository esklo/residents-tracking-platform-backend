package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/request"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	requestRepository "github.com/esklo/residents-tracking-platform-backend/internal/repository/request"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	requestService "github.com/esklo/residents-tracking-platform-backend/internal/service/request"
)

func (s *ServiceProvider) RequestRepository() repository.RequestRepository {
	if s.requestRepository == nil {
		s.requestRepository = requestRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.requestRepository
}

func (s *ServiceProvider) RequestService() service.RequestService {
	if s.requestService == nil {
		s.requestService = requestService.NewService(
			s.RequestRepository(),
			s.ContactService(),
			s.FileService(),
		)
	}

	return s.requestService
}

func (s *ServiceProvider) RequestImpl() *request.Implementation {
	if s.requestImpl == nil {
		s.requestImpl = request.NewImplementation(s.RequestService(), s.AuthService())
	}

	return s.requestImpl
}
