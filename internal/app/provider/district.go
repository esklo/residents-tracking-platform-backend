package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/api/district"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	districtRepository "github.com/esklo/residents-tracking-platform-backend/internal/repository/district"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	districtService "github.com/esklo/residents-tracking-platform-backend/internal/service/district"
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
			s.GetLogger(),
		)
	}

	return s.districtService
}

func (s *ServiceProvider) DistrictImpl() *district.Implementation {
	if s.districtImpl == nil {
		s.districtImpl = district.NewImplementation(s.DistrictService(), s.AuthService(), s.GetLogger())
	}

	return s.districtImpl
}
