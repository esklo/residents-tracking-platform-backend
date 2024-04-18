package provider

import (
	"github.com/esklo/residents-tracking-platform/internal/api/contact"
	"github.com/esklo/residents-tracking-platform/internal/repository"
	contactRepository "github.com/esklo/residents-tracking-platform/internal/repository/contact"
	"github.com/esklo/residents-tracking-platform/internal/service"
	contactService "github.com/esklo/residents-tracking-platform/internal/service/contact"
)

func (s *ServiceProvider) ContactRepository() repository.ContactRepository {
	if s.contactRepository == nil {
		s.contactRepository = contactRepository.NewRepository(
			s.DatabaseConnection,
		)
	}

	return s.contactRepository
}

func (s *ServiceProvider) ContactService() service.ContactService {
	if s.contactService == nil {
		s.contactService = contactService.NewService(
			s.ContactRepository(),
		)
	}

	return s.contactService
}

func (s *ServiceProvider) ContactImpl() *contact.Implementation {
	if s.contactImpl == nil {
		s.contactImpl = contact.NewImplementation(s.ContactService(), s.AuthService())
	}

	return s.contactImpl
}
