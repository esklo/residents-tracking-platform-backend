package contact

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform/gen/proto/contact"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/esklo/residents-tracking-platform/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Implementation struct {
	proto.UnimplementedContactServiceServer
	contactService service.ContactService
	authService    service.AuthService
}

func NewImplementation(contactService service.ContactService, authService service.AuthService) *Implementation {
	return &Implementation{
		contactService: contactService,
		authService:    authService,
	}
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Contact, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	contactId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse contact id")
	}
	contact, err := i.contactService.Get(ctx, &contactId)
	if err != nil {
		return nil, err
	}
	return contact.ToProto()
}
