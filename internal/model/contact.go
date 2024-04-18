package model

import (
	protoContact "github.com/esklo/residents-tracking-platform/gen/proto/contact"
	"github.com/google/uuid"
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
