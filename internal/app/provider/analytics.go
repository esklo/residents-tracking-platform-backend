package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/analytics"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	analyticsService "github.com/esklo/residents-tracking-platform-backend/internal/service/analytics"
)

func (s *ServiceProvider) AnalyticsService() service.AnalyticsService {
	if s.analyticsService == nil {
		s.analyticsService = analyticsService.NewService(
			s.ThemeService(),
			s.RequestService(),
			s.DepartmentService(),
			s.GetLogger(),
		)
	}

	return s.analyticsService
}

func (s *ServiceProvider) AnalyticsImpl() *analytics.Implementation {
	if s.analyticsImpl == nil {
		s.analyticsImpl = analytics.NewImplementation(s.AuthService(), s.AnalyticsService(), s.GetLogger())
	}

	return s.analyticsImpl
}
