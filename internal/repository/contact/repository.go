package contact

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform/internal/model"
	def "github.com/esklo/residents-tracking-platform/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var _ def.ContactRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r Repository) GetByID(_ context.Context, id string) (*model.Contact, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "contact get by id: can not get database connection")
	}

	var contact model.Contact
	err = connection.QueryRow(`select id, phone, email, note, name from contacts where id = $1`, id).
		Scan(
			&contact.Id,
			&contact.Phone,
			&contact.Email,
			&contact.Note,
			&contact.Name,
		)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r Repository) Create(ctx context.Context, contact *model.Contact) (*model.Contact, error) {
	if contact.Id == uuid.Nil {
		contact.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into contacts (id,phone, email, note, name) 
		values ($1,$2,$3,$4,$5)
	`, contact.Id, contact.Phone, contact.Email, contact.Note, contact.Name)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return contact, nil
}

func (r Repository) Update(ctx context.Context, contact *model.Contact) error {
	if contact == nil {
		return errors.New("contact is nil")
	}

	if contact.Id == uuid.Nil {
		return errors.New("contact id is nil")
	}

	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		update contacts set phone = $1, email = $2,note=$3,name=$4 where id = $5
	`, contact.Phone, contact.Email, contact.Note, contact.Name, contact.Id)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}
