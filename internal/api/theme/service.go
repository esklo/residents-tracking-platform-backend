package theme

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/theme"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Implementation struct {
	proto.UnimplementedThemeServiceServer
	themeService      service.ThemeService
	departmentService service.DepartmentService
	authService       service.AuthService
}

func NewImplementation(themeService service.ThemeService, departmentService service.DepartmentService, authService service.AuthService) *Implementation {
	return &Implementation{
		themeService:      themeService,
		departmentService: departmentService,
		authService:       authService,
	}
}

func (i Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.Theme, error) {
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	departmentId, err := uuid.Parse(req.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse department id")
	}

	if !user.IsAdmin() {
		if user.DepartmentId == nil {
			return nil, errors.New("user not linked to department")
		}
		userDepartment, err := i.departmentService.Get(ctx, user.DepartmentId)
		if err != nil {
			return nil, err
		}

		reqDepartment, err := i.departmentService.Get(ctx, &departmentId)
		if err != nil {
			return nil, err
		}

		if userDepartment.DistrictId != reqDepartment.DistrictId {
			return nil, model.ErrorPermissionDenied
		}
	}
	theme, err := i.themeService.Create(ctx, &model.Theme{
		Title:        req.Title,
		Priority:     req.Priority,
		DepartmentId: departmentId,
	})
	if err != nil {
		return nil, err
	}
	return theme.ToProto()
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Theme, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	themeId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse theme id")
	}
	theme, err := i.themeService.Get(ctx, &themeId)
	if err != nil {
		return nil, err
	}
	return theme.ToProto()
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

	themes, err := i.themeService.GetAll(ctx, districtId)
	if err != nil {
		return nil, err
	}
	var protoThemes []*proto.Theme
	for _, theme := range themes {
		protoTheme, err := theme.ToProto()
		if err != nil {
			return nil, err
		}
		protoThemes = append(protoThemes, protoTheme)
	}
	return &proto.GetResponse{
		Themes: protoThemes,
	}, nil
}

func (i Implementation) Update(ctx context.Context, req *proto.Theme) (*proto.Theme, error) {
	themeId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse theme id")
	}
	departmentId, err := uuid.Parse(req.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse department id")
	}
	theme := model.Theme{
		Id:           themeId,
		Title:        req.Title,
		Priority:     req.Priority,
		DepartmentId: departmentId,
	}
	if err := i.themeService.Update(ctx, &theme); err != nil {
		return nil, err
	}
	return theme.ToProto()
}

func (i Implementation) Delete(ctx context.Context, req *proto.ByIdRequest) (*empty.Empty, error) {
	themeId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse theme id")
	}
	if err := i.themeService.Delete(ctx, &themeId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
