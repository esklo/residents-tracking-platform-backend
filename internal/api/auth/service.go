package auth

import (
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/auth"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"go.uber.org/zap"
)

type Implementation struct {
	proto.UnimplementedAuthServiceServer
	authService service.AuthService
	logger      *zap.Logger
}

func NewImplementation(authService service.AuthService, logger *zap.Logger) *Implementation {
	return &Implementation{
		authService: authService,
		logger:      logger,
	}
}
