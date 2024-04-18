package user

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i *Implementation) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	var districtId *uuid.UUID
	if !user.IsAdmin() {
		if user.DepartmentId == nil {
			return nil, errors.New("user not linked to department")
		}
		department, err := i.departmentService.Get(ctx, user.DepartmentId)
		if err != nil {
			return nil, err
		}
		districtId = &department.DistrictId
	}

	users, err := i.userService.GetAll(ctx, districtId)
	if err != nil {
		return nil, err
	}
	var protoUsers []*proto.User
	for _, user := range users {
		protoUser, err := user.ToProto()
		if err != nil {
			return nil, err
		}
		protoUsers = append(protoUsers, protoUser)
	}
	return &proto.GetResponse{
		Users: protoUsers,
	}, nil
}
