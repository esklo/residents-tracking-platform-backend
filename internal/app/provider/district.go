package provider

import (
	"github.com/esklo/residents-tracking-platform/internal/api/district"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	districtRepository "github.com/esklo/residents-tracking-platform/internal/repository/district"
	"github.com/esklo/residents-tracking-platform/internal/service"
	districtService "github.com/esklo/residents-tracking-platform/internal/service/district"
)

func (s *ServiceProvider) DistrictRepository() repository.DistrictRepository {
	if s.districtRepository == nil {
		s.districtRepository = districtRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.districtRepository
}

func (s *ServiceProvider) DistrictService() service.DistrictService {
	if s.districtService == nil {
		s.districtService = districtService.NewService(
			s.DistrictRepository(),
			s.FileRepository(),
		)
	}

	return s.districtService
}

func (s *ServiceProvider) DistrictImpl() *district.Implementation {
	if s.districtImpl == nil {
		s.districtImpl = district.NewImplementation(s.DistrictService(), s.AuthService())
	}

	return s.districtImpl
}
