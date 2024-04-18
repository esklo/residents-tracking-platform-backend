package user

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
)

var _ def.UserService = (*Service)(nil)

type Service struct {
	userRepository    repository.UserRepository
	departmentService def.DepartmentService
}

func NewService(
	userRepository repository.UserRepository,
	departmentService def.DepartmentService,
) *Service {
	return &Service{
		userRepository:    userRepository,
		departmentService: departmentService,
	}
}
