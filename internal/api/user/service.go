package user

import (
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
)

type Implementation struct {
	proto.UnimplementedUserServiceServer
	userService       service.UserService
	authService       service.AuthService
	departmentService service.DepartmentService
}

func NewImplementation(userService service.UserService, authService service.AuthService, departmentService service.DepartmentService) *Implementation {
	return &Implementation{
		userService:       userService,
		authService:       authService,
		departmentService: departmentService,
	}
}
