package auth

import (
	proto "github.com/esklo/residents-tracking-platform/gen/proto/auth"
	"github.com/esklo/residents-tracking-platform/internal/service"
)

type Implementation struct {
	proto.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
