package auth

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/auth"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/pkg/errors"
)

func (i Implementation) ChangePassword(ctx context.Context, req *proto.ChangePasswordRequest) (*empty.Empty, error) {
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	if !user.CheckPassword(req.CurrentPassword) {
		return nil, model.ErrorCurrentPasswordIsInvalid
	}

	if req.NewPassword != req.NewPasswordConfirmation {
		return nil, errors.New("new passwords mismatch")
	}

	if err := i.authService.ChangePassword(ctx, user, req.NewPassword); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
