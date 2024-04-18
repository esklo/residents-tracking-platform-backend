package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform/internal/model"
)

func (s *Service) Update(ctx context.Context, user *model.User) error {
	return s.userRepository.Update(ctx, user)
}
