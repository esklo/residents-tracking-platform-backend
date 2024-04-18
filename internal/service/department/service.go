package department

import (
	"context"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	def "github.com/esklo/residents-tracking-platform/internal/service"
	"github.com/google/uuid"
)

var _ def.DepartmentService = (*Service)(nil)

type Service struct {
	departmentRepository repository.DepartmentRepository
}

func NewService(
	departmentRepository repository.DepartmentRepository,
) *Service {
	return &Service{
		departmentRepository: departmentRepository,
	}
}

func (s *Service) Create(ctx context.Context, department *model.Department) (*model.Department, error) {
	return s.departmentRepository.Create(ctx, department)
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.Department, error) {
	return s.departmentRepository.GetByID(ctx, id.String())
}

func (s *Service) GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.Department, error) {
	if districtId == nil {
		return s.departmentRepository.GetAll(ctx)
	}
	return s.departmentRepository.GetAllWithDistrictId(ctx, districtId.String())
}
