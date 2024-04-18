package user

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i *Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.User, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	departmentId, err := uuid.Parse(req.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse department id")
	}
	user := model.User{
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		FatherName:   req.FatherName,
		Password:     req.Password,
		DepartmentId: &departmentId,
	}
	switch req.Role {
	case proto.Role_EmployeeRole:
		user.Role = model.UserRoleEmployee
	default:
		return nil, errors.New("role is invalid")
	}
	createdUser, err := i.userService.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return createdUser.ToProto()
}
