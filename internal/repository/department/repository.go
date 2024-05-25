package department

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	def "github.com/esklo/residents-tracking-platform-backend/internal/repository"
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
	err = connection.QueryRow(`select id, title, district_id, full_access from departments where id = $1`, id).
		Scan(
			&department.Id,
			&department.Title,
			&department.DistrictId,
			&department.FullAccess,
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
		insert into departments (title,district_id,full_access) 
		values ($1,$2,$3)
	`, department.Title, department.DistrictId, department.FullAccess)
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
	rows, err := connection.Query("select id, title, district_id, full_access from departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(&department.Id, &department.Title, &department.DistrictId, &department.FullAccess)
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
	rows, err := connection.Query("select id, title, district_id, full_access from departments where district_id=$1", districtId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*model.Department
	for rows.Next() {
		var department model.Department
		err := rows.Scan(&department.Id, &department.Title, &department.DistrictId, &department.FullAccess)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}
	return departments, nil
}

func (r *Repository) Update(ctx context.Context, department *model.Department) error {
	if department == nil {
		return errors.New("department is nil")
	}

	if department.Id == uuid.Nil {
		return errors.New("department id is nil")
	}

	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		update departments set title = $1, district_id = $2, full_access=$3 where id = $4
	`, department.Title, department.DistrictId, department.FullAccess, department.Id)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}
