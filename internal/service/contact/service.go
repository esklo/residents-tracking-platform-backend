package contact

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
)

var _ def.ContactService = (*Service)(nil)

type Service struct {
	contactRepository repository.ContactRepository
}

func NewService(
	contactRepository repository.ContactRepository,
) *Service {
	return &Service{
		contactRepository: contactRepository,
	}
}

func (s *Service) Create(ctx context.Context, contact *model.Contact) (*model.Contact, error) {
	return s.contactRepository.Create(ctx, contact)
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.Contact, error) {
	return s.contactRepository.GetByID(ctx, id.String())
}
