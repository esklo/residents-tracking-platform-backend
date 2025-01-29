package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) Get(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("user get error", zap.Error(err))
		return nil, err
	}
	if user == nil {
		s.logger.Error("user with id not found", zap.String("id", id))
		return nil, model.ErrorNotFound
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.User, error) {
	if districtId == nil {
		return s.userRepository.GetAll(ctx)
	}
	var departmentIds []string
	departments, err := s.departmentService.GetAll(ctx, districtId)
	if err != nil {
		return nil, err
	}
	for _, department := range departments {
		departmentIds = append(departmentIds, department.Id.String())
	}
	return s.userRepository.GetAllWithDepartmentIds(ctx, departmentIds)
}

func (s *Service) GetAllWithinDepartment(ctx context.Context, departmentId *uuid.UUID) ([]*model.User, error) {
	if departmentId == nil {
		return nil, nil
	}
	return s.userRepository.GetAllWithDepartmentIds(ctx, []string{departmentId.String()})
}

func (s *Service) GetAllForThemeId(ctx context.Context, themeId *uuid.UUID) ([]*model.User, error) {
	if themeId == nil {
		return nil, nil
	}
	theme, err := s.themeService.Get(ctx, themeId)
	if err != nil {
		return nil, err
	}
	return s.GetAllWithinDepartment(ctx, &theme.DepartmentId)
}
