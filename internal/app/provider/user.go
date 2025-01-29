package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	userRepository "github.com/esklo/residents-tracking-platform-backend/internal/repository/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	userService "github.com/esklo/residents-tracking-platform-backend/internal/service/user"
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
			s.ThemeService(),
			s.GetLogger(),
		)
	}

	return s.userService
}

func (s *ServiceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(), s.AuthService(), s.DepartmentService(), s.GetLogger())
	}

	return s.userImpl
}
