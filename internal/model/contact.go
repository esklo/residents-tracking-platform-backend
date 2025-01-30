package model

import (
	protoContact "github.com/esklo/residents-tracking-platform-backend/gen/proto/contact"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Contact struct {
	Id    uuid.UUID
	Phone *int64
	Email *string
	Name  string
	Note  *string
}

func (c *Contact) ToProto() (*protoContact.Contact, error) {
	if c == nil {
		return nil, ErrorModelIsEmpty
	}
	contact := protoContact.Contact{
		Id:    c.Id.String(),
		Phone: c.Phone,
		Email: c.Email,
		Name:  c.Name,
		Note:  c.Note,
	}
	return &contact, nil
}

func (c *Contact) FromProto(contact *protoContact.Contact) error {
	if contact == nil {
		return nil
	}
	contactId, err := uuid.Parse(contact.Id)
	if err != nil {
		return errors.Wrap(err, "can not parse contact id")
	}
	c.Id = contactId
	c.Phone = contact.Phone
	c.Email = contact.Email
	c.Name = contact.Name
	c.Note = contact.Note
	return nil
}
