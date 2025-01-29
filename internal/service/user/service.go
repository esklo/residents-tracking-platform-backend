package user

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"go.uber.org/zap"
)

var _ def.UserService = (*Service)(nil)

type Service struct {
	userRepository    repository.UserRepository
	departmentService def.DepartmentService
	themeService      def.ThemeService
	logger            *zap.Logger
}

func NewService(
	userRepository repository.UserRepository,
	departmentService def.DepartmentService,
	themeService def.ThemeService,
	logger *zap.Logger,
) *Service {
	return &Service{
		userRepository:    userRepository,
		departmentService: departmentService,
		themeService:      themeService,
		logger:            logger,
	}
}
