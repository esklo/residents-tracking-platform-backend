package user

import (
	"context"
	protoEmpty "github.com/esklo/residents-tracking-platform/gen/proto/empty"
	protoUser "github.com/esklo/residents-tracking-platform/gen/proto/user"
	"github.com/esklo/residents-tracking-platform/internal/model"
)

func (i *Implementation) GetMe(ctx context.Context, _ *protoEmpty.Empty) (*protoUser.User, error) {
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}
	proto, err := user.ToProto()
	if err != nil {
		return nil, err
	}
	return proto, nil
}
