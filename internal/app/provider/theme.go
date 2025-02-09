package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/theme"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	themeRepository "github.com/esklo/residents-tracking-platform-backend/internal/repository/theme"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	themeService "github.com/esklo/residents-tracking-platform-backend/internal/service/theme"
)

func (s *ServiceProvider) ThemeRepository() repository.ThemeRepository {
	if s.themeRepository == nil {
		s.themeRepository = themeRepository.NewRepository(s.DatabaseConnection)
	}

	return s.themeRepository
}

func (s *ServiceProvider) ThemeService() service.ThemeService {
	if s.themeService == nil {
		s.themeService = themeService.NewService(
			s.ThemeRepository(),
			s.DepartmentService(),
			s.GetLogger(),
		)
	}

	return s.themeService
}

func (s *ServiceProvider) ThemeImpl() *theme.Implementation {
	if s.themeImpl == nil {
		s.themeImpl = theme.NewImplementation(s.ThemeService(), s.DepartmentService(), s.AuthService(), s.GetLogger())
	}

	return s.themeImpl
}
