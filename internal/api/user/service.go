package user

import (
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"go.uber.org/zap"
)

type Implementation struct {
	proto.UnimplementedUserServiceServer
	userService       service.UserService
	authService       service.AuthService
	departmentService service.DepartmentService
	logger            *zap.Logger
}

func NewImplementation(userService service.UserService, authService service.AuthService, departmentService service.DepartmentService, logger *zap.Logger) *Implementation {
	return &Implementation{
		userService:       userService,
		authService:       authService,
		departmentService: departmentService,
		logger:            logger,
	}
}
