package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/geo"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	geoService "github.com/esklo/residents-tracking-platform-backend/internal/service/geo"
)

func (s *ServiceProvider) GeoService() service.GeoService {
	if s.geoService == nil {
		s.geoService = geoService.NewService(s.DadataClient())
	}

	return s.geoService
}

func (s *ServiceProvider) GeoImpl() *geo.Implementation {
	if s.geoImpl == nil {
		s.geoImpl = geo.NewImplementation(s.GeoService(), s.AuthService())
	}

	return s.geoImpl
}
