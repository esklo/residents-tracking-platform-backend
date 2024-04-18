package provider

import (
	"github.com/esklo/residents-tracking-platform/internal/api/theme"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	themeRepository "github.com/esklo/residents-tracking-platform/internal/repository/theme"
	"github.com/esklo/residents-tracking-platform/internal/service"
	themeService "github.com/esklo/residents-tracking-platform/internal/service/theme"
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
		)
	}

	return s.themeService
}

func (s *ServiceProvider) ThemeImpl() *theme.Implementation {
	if s.themeImpl == nil {
		s.themeImpl = theme.NewImplementation(s.ThemeService(), s.DepartmentService(), s.AuthService())
	}

	return s.themeImpl
}
