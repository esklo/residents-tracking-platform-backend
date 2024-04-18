package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"log"
)

func (s *Service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := user.SetPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user, err = s.userRepository.Create(ctx, user)
	if err != nil {
		log.Printf("ошибка создания пользователя: %v\n", err)
		return nil, err
	}

	return user, nil
}
