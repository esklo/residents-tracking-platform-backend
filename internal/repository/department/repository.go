package department

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform/internal/model"
	def "github.com/esklo/residents-tracking-platform/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var _ def.DepartmentRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(_ context.Context, id string) (*model.Department, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "department get by id: can not get database connection")
	}

	var department model.Department
	err = connection.QueryRow(`select id, title, district_id from departments where id = $1`, id).
		Scan(
			&department.Id,
			&department.Title,
			&department.DistrictId,
		)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *Repository) Create(_ context.Context, department *model.Department) (*model.Department, error) {
	if department.Id == uuid.Nil {
		department.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into departments (title,district_id) 
		values ($1,$2)
	`, department.Title, department.DistrictId)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return department, nil
}

func (r *Repository) GetAll(_ context.Context) ([]*model.Department, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select id, title, district_id from departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(&department.Id, &department.Title, &department.DistrictId)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}
	return departments, nil
}

func (r *Repository) GetAllWithDistrictId(ctx context.Context, districtId string) ([]*model.Department, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select id, title, district_id from departments where district_id=$1", districtId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(&department.Id, &department.Title, &department.DistrictId)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}
	return departments, nil
}
