package department

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/department"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Implementation struct {
	proto.UnimplementedDepartmentServiceServer
	departmentService service.DepartmentService
	authService       service.AuthService
}

func NewImplementation(departmentService service.DepartmentService, authService service.AuthService) *Implementation {
	return &Implementation{
		departmentService: departmentService,
		authService:       authService,
	}
}

func (i Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.Department, error) {
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	districtId, err := uuid.Parse(req.DistrictId)
	if !user.IsAdmin() {
		if user.DepartmentId == nil {
			return nil, errors.New("user not linked to department")
		}
		userDepartment, err := i.departmentService.Get(ctx, user.DepartmentId)
		if err != nil {
			return nil, err
		}

		if districtId != userDepartment.DistrictId {
			return nil, model.ErrorPermissionDenied
		}
	}

	if err != nil {
		return nil, errors.Wrap(err, "can not parse district id")
	}
	department, err := i.departmentService.Create(ctx, &model.Department{
		Title:      req.Title,
		DistrictId: districtId,
	})
	if err != nil {
		return nil, err
	}
	return department.ToProto()
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Department, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	departmentId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse department id")
	}
	department, err := i.departmentService.Get(ctx, &departmentId)
	if err != nil {
		return nil, err
	}
	return department.ToProto()
}

func (i Implementation) Get(ctx context.Context, _ *proto.GetRequest) (*proto.GetResponse, error) {
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

	departments, err := i.departmentService.GetAll(ctx, districtId)
	if err != nil {
		return nil, err
	}
	var protoDepartments []*proto.Department
	for _, department := range departments {
		protoDepartment, err := department.ToProto()
		if err != nil {
			return nil, err
		}
		protoDepartments = append(protoDepartments, protoDepartment)
	}
	return &proto.GetResponse{
		Departments: protoDepartments,
	}, nil
}
