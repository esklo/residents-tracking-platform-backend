package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"log"
)

func (s *Service) Get(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		log.Printf("ошибка получения пользователя: %v\n", err)
		return nil, err
	}
	if user == nil {
		log.Printf("пользователь с id %s не найден\n", id)
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
