package user

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i *Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.User, error) {
	i.logger.Debug("user.GetById request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse user id")
	}

	user, err := i.userService.Get(ctx, &userId)
	if err != nil {
		return nil, err
	}

	return user.ToProto()
}
