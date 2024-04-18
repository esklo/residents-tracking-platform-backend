package district

import (
	"context"
	"database/sql"
	"github.com/esklo/residents-tracking-platform/internal/model"
	def "github.com/esklo/residents-tracking-platform/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var _ def.DistrictRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(_ context.Context, id string) (*model.District, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "district get by id: can not get database connection")
	}

	var district model.District
	err = connection.QueryRow(`select id, title,geojson,coat_of_arms_file_id from districts where id = $1`, id).
		Scan(
			&district.Id,
			&district.Title,
			&district.GeoJson,
			&district.CoatOfArmsFileId,
		)
	if err != nil {
		return nil, err
	}

	return &district, nil
}

func (r *Repository) Create(_ context.Context, district *model.District) (*model.District, error) {
	if district.Id == uuid.Nil {
		district.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into districts (title,geojson,coat_of_arms_file_id) 
		values ($1,$2,$3)
	`, district.Title, district.GeoJson, district.CoatOfArmsFileId)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return district, nil
}

func (r *Repository) GetAll(_ context.Context) ([]*model.District, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select id, title,geojson,coat_of_arms_file_id from districts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var districts []*model.District
	for rows.Next() {
		var district model.District
		err := rows.Scan(&district.Id, &district.Title, &district.GeoJson, &district.CoatOfArmsFileId)
		if err != nil {
			return nil, err
		}
		districts = append(districts, &district)
	}
	return districts, nil
}
