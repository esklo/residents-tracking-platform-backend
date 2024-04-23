package theme

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	def "github.com/esklo/residents-tracking-platform-backend/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
)

var _ def.ThemeRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(ctx context.Context, id string) (*model.Theme, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "theme get by id: can not get database connection")
	}

	var theme model.Theme
	err = connection.QueryRow(`select id, title, priority, department_id, deleted_at from themes where id = $1`, id).
		Scan(
			&theme.Id,
			&theme.Title,
			&theme.Priority,
			&theme.DepartmentId,
			&theme.DeletedAt,
		)
	if err != nil {
		return nil, err
	}

	return &theme, nil
}

func (r *Repository) Create(ctx context.Context, theme *model.Theme) (*model.Theme, error) {
	if theme.Id == uuid.Nil {
		theme.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into themes (title,priority,department_id) 
		values ($1,$2,$3)
	`, theme.Title, theme.Priority, theme.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return theme, nil
}

func (r *Repository) GetAll(ctx context.Context, onlyDeleted bool) ([]*model.Theme, error) {
	deletedText := ""
	if onlyDeleted {
		deletedText = "not"
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query(fmt.Sprintf("select id, title, priority, department_id, deleted_at from themes where deleted_at is %s null order by id", deletedText))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var themes []*model.Theme
	for rows.Next() {
		var theme model.Theme
		err := rows.Scan(&theme.Id, &theme.Title, &theme.Priority, &theme.DepartmentId, &theme.DeletedAt)
		if err != nil {
			return nil, err
		}
		themes = append(themes, &theme)
	}
	return themes, nil
}

func (r *Repository) GetAllWithDepartmentIds(ctx context.Context, departmentIds []string, onlyDeleted bool) ([]*model.Theme, error) {
	deletedText := "is"
	if onlyDeleted {
		deletedText = "is not"
	}

	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query(
		fmt.Sprintf("select id, title, priority, department_id, deleted_at from themes where department_id in (%s) and deleted_at %s null order by id",
			fmt.Sprintf("'%s'", strings.Join(departmentIds, "','")),
			deletedText,
		),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var themes []*model.Theme
	for rows.Next() {
		var theme model.Theme
		err := rows.Scan(&theme.Id, &theme.Title, &theme.Priority, &theme.DepartmentId, &theme.DeletedAt)
		if err != nil {
			return nil, err
		}
		themes = append(themes, &theme)
	}
	return themes, nil
}

func (r *Repository) Update(ctx context.Context, theme *model.Theme) error {
	if theme == nil {
		return errors.New("theme is nil")
	}

	if theme.Id == uuid.Nil {
		return errors.New("theme id is nil")
	}

	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		update themes set title = $1, priority = $2, department_id=$3, deleted_at=$4 where id = $5
	`, theme.Title, theme.Priority, theme.DepartmentId, theme.DeletedAt, theme.Id)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}
