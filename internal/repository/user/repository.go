package user

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

var _ def.UserRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(_ context.Context, id string) (*model.User, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "user get by id: can not get database connection")
	}

	var user model.User
	err = connection.QueryRow(`select id, email, role,first_name, last_name, father_name, password, salt, department_id from users where id = $1`, id).
		Scan(
			&user.Id,
			&user.Email,
			&user.Role,
			&user.FirstName,
			&user.LastName,
			&user.FatherName,
			&user.Password,
			&user.Salt,
			&user.DepartmentId,
		)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Create(_ context.Context, user *model.User) (*model.User, error) {
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	_, err = connection.Exec(`
		insert into users (email, role, first_name, last_name, father_name, password, salt, department_id) 
		values ($1,$2,$3,$4,$5,$6,$7,$8)
	`, user.Email, user.Role, user.FirstName, user.LastName, user.FatherName, user.Password, user.Salt, user.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return user, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "user get by email: can not get database connection")
	}

	var user model.User
	err = connection.QueryRow(`select id, email, role, first_name, last_name, father_name, password, salt, department_id from users where email = $1`, email).
		Scan(
			&user.Id,
			&user.Email,
			&user.Role,
			&user.FirstName,
			&user.LastName,
			&user.FatherName,
			&user.Password,
			&user.Salt,
			&user.DepartmentId,
		)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *Repository) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.Id == uuid.Nil {
		return errors.New("user id is nil")
	}

	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		update users set email = $1, role = $2,first_name=$3,last_name=$4,father_name=$5,password=$6,salt=$7,department_id=$8 where id = $9
	`, user.Email, user.Role, user.FirstName, user.LastName, user.FatherName, user.Password, user.Salt, user.DepartmentId, user.Id)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*model.User, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select id, email,first_name,last_name,father_name,department_id from users order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.FatherName, &user.DepartmentId)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *Repository) GetAllWithDepartmentIds(ctx context.Context, departmentIds []string) ([]*model.User, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query(
		fmt.Sprintf("select id, email, first_name, last_name, father_name, department_id from users where department_id in (%s) order by id",
			fmt.Sprintf("'%s'", strings.Join(departmentIds, "','"))),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.FatherName, &user.DepartmentId)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
