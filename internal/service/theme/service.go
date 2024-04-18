package theme

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"time"
)

var _ def.ThemeService = (*Service)(nil)

type Service struct {
	themeRepository   repository.ThemeRepository
	departmentService def.DepartmentService
}

func NewService(
	themeRepository repository.ThemeRepository,
	departmentService def.DepartmentService,
) *Service {
	return &Service{
		themeRepository:   themeRepository,
		departmentService: departmentService,
	}
}

func (s *Service) Create(ctx context.Context, department *model.Theme) (*model.Theme, error) {
	return s.themeRepository.Create(ctx, department)
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.Theme, error) {
	return s.themeRepository.GetByID(ctx, id.String())
}

func (s *Service) GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.Theme, error) {
	if districtId == nil {
		return s.themeRepository.GetAll(ctx, false)
	}
	var departmentIds []string
	departments, err := s.departmentService.GetAll(ctx, districtId)
	if err != nil {
		return nil, err
	}
	for _, department := range departments {
		departmentIds = append(departmentIds, department.Id.String())
	}
	return s.themeRepository.GetAllWithDepartmentIds(ctx, departmentIds, false)
}

func (s *Service) Update(ctx context.Context, theme *model.Theme) error {
	return s.themeRepository.Update(ctx, theme)
}

func (s *Service) Delete(ctx context.Context, id *uuid.UUID) error {
	theme, err := s.Get(ctx, id)
	if err != nil {
		return err
	}
	deletedAt := time.Now()
	theme.DeletedAt = &deletedAt
	return s.Update(ctx, theme)
}
