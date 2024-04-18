package user

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i *Implementation) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.User, error) {
	currentUser, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if currentUser.Role != model.UserRoleAdmin && currentUser.Id.String() != req.Id {
		return nil, model.ErrorPermissionDenied
	}

	updatedUser := user

	if currentUser.Role == model.UserRoleAdmin {
		if req.Email != nil {
			updatedUser.Email = *req.Email
		}
		if req.Role != nil {
			switch *req.Role {
			case proto.Role_AdminRole:
				updatedUser.Role = model.UserRoleAdmin
			}
		}
		if req.DepartmentId != nil {
			departmentId, err := uuid.Parse(*req.DepartmentId)
			if err != nil {
				return nil, errors.Wrap(err, "can not parse department id")
			}
			updatedUser.DepartmentId = &departmentId
		}
		if req.Password != nil {
			err := updatedUser.SetPassword(*req.Password)
			if err != nil {
				return nil, err
			}
		}
	}

	updatedUser.LastName = req.LastName
	if req.FirstName != nil {
		updatedUser.FirstName = *req.FirstName
	}
	updatedUser.FatherName = req.FatherName

	if err := i.userService.Update(ctx, updatedUser); err != nil {
		return nil, err
	}

	updatedUserProto, err := updatedUser.ToProto()
	if err != nil {
		return nil, err
	}
	return updatedUserProto, nil
}
