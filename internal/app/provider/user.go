package provider

import (
	"github.com/esklo/residents-tracking-platform/internal/api/user"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	userRepository "github.com/esklo/residents-tracking-platform/internal/repository/user"
	"github.com/esklo/residents-tracking-platform/internal/service"
	userService "github.com/esklo/residents-tracking-platform/internal/service/user"
)

func (s *ServiceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.userRepository
}

func (s *ServiceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
			s.DepartmentService(),
		)
	}

	return s.userService
}

func (s *ServiceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(), s.AuthService(), s.DepartmentService())
	}

	return s.userImpl
}