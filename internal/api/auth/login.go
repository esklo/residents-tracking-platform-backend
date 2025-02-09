package auth

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/auth"
)

func (i Implementation) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	i.logger.Debug("auth.Login request")
	user, err := i.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := i.authService.CreateToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	protoUser, err := user.ToProto()
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token: token,
		User:  protoUser,
	}, nil
}
