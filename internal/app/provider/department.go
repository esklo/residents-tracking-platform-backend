package provider

import (
	"github.com/esklo/residents-tracking-platform/internal/api/department"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	departmentRepository "github.com/esklo/residents-tracking-platform/internal/repository/department"
	"github.com/esklo/residents-tracking-platform/internal/service"
	departmentService "github.com/esklo/residents-tracking-platform/internal/service/department"
)

func (s *ServiceProvider) DepartmentRepository() repository.DepartmentRepository {
	if s.departmentRepository == nil {
		s.departmentRepository = departmentRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.departmentRepository
}

func (s *ServiceProvider) DepartmentService() service.DepartmentService {
	if s.departmentService == nil {
		s.departmentService = departmentService.NewService(
			s.DepartmentRepository(),
		)
	}

	return s.departmentService
}

func (s *ServiceProvider) DepartmentImpl() *department.Implementation {
	if s.departmentImpl == nil {
		s.departmentImpl = department.NewImplementation(s.DepartmentService(), s.AuthService())
	}

	return s.departmentImpl
}
