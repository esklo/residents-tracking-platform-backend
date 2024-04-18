package user

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
)

func (i *Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.User, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return user.ToProto()
}
