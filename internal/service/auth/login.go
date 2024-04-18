package auth

import (
	"context"
	"github.com/esklo/residents-tracking-platform/internal/model"
)

func (s *Service) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, model.ErrorInvalidCredentials
	}
	if !user.CheckPassword(password) {
		return nil, model.ErrorInvalidCredentials
	}
	return user, nil
}

func (s *Service) ChangePassword(ctx context.Context, user *model.User, password string) error {
	if err := user.SetPassword(password); err != nil {
		return err
	}
	if err := s.userRepository.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
