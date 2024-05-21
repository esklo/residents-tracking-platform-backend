package department

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ def.DepartmentService = (*Service)(nil)

type Service struct {
	departmentRepository repository.DepartmentRepository
	logger               *zap.Logger
}

func NewService(
	departmentRepository repository.DepartmentRepository,
	logger *zap.Logger,
) *Service {
	return &Service{
		departmentRepository: departmentRepository,
		logger:               logger,
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
