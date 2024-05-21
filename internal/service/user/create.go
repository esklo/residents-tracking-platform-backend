package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := user.SetPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user, err = s.userRepository.Create(ctx, user)
	if err != nil {
		s.logger.Error("user creation error", zap.Error(err))
		return nil, err
	}

	return user, nil
}
