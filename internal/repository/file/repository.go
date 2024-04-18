package file

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform/internal/model"
	def "github.com/esklo/residents-tracking-platform/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var _ def.FileRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(ctx context.Context, id string) (*model.File, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "file get by id: can not get database connection")
	}

	var file model.File
	err = connection.QueryRow(`select id, filename, mimetype, extension, path from files where id = $1`, id).
		Scan(
			&file.Id,
			&file.FileName,
			&file.MimeType,
			&file.Extension,
			&file.Path,
		)
	if err != nil {
		return nil, errors.Wrap(err, "can not query row (file)")
	}

	return &file, nil
}

func (r *Repository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	if file.Id == uuid.Nil {
		file.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into files (id, filename, mimetype, extension, path) 
		values ($1,$2,$3,$4,$5)
	`, file.Id, file.FileName, file.MimeType, file.Extension, file.Path)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return file, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}
	_, err = connection.Exec(`delete from files where id = $1`, id)
	if err != nil {
		return errors.Wrap(err, "can not delete file by id")
	}
	return nil
}
