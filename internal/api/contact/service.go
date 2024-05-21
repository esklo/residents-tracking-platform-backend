package contact

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/contact"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Implementation struct {
	proto.UnimplementedContactServiceServer
	contactService service.ContactService
	authService    service.AuthService
	logger         *zap.Logger
}

func NewImplementation(contactService service.ContactService, authService service.AuthService, logger *zap.Logger) *Implementation {
	return &Implementation{
		contactService: contactService,
		authService:    authService,
		logger:         logger,
	}
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Contact, error) {
	i.logger.Debug("contact.GetById request")
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
