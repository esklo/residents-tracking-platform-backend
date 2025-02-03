package request

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	def "github.com/esklo/residents-tracking-platform-backend/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"strings"
	"time"
)

var _ def.RequestRepository = (*Repository)(nil)

type Repository struct {
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		getConnection: getConnection,
	}
}

func (r *Repository) GetByID(ctx context.Context, id *uuid.UUID) (*model.Request, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "request get by id: can not get database connection")
	}

	var geo []byte
	var request model.Request
	var contactId uuid.UUID
	err = connection.QueryRow(`select id, description, ST_AsEWKB(geo), address, created_at, deleted_at, status, priority, theme_id, user_id, contact_id, deadline, comment from requests where id = $1`, id).
		Scan(
			&request.Id,
			&request.Description,
			&geo,
			&request.Address,
			&request.CreatedAt,
			&request.DeletedAt,
			&request.Status,
			&request.Priority,
			&request.ThemeId,
			&request.UserId,
			&contactId,
			&request.Deadline,
			&request.Comment,
		)
	if err != nil {
		return nil, err
	}
	point, err := ewkb.Unmarshal(geo)
	if err != nil {
		return nil, err
	}
	request.Geo = model.GeoPoint{
		Lat: point.FlatCoords()[1],
		Lon: point.FlatCoords()[0],
	}

	request.Contact = &model.Contact{
		Id: contactId,
	}
	return &request, nil
}

func (r *Repository) Create(ctx context.Context, request *model.Request) (*model.Request, error) {
	if request.Id == uuid.Nil {
		request.Id = uuid.New()
	}
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into requests (id, description, geo, address, created_at, deleted_at, status, priority, theme_id, user_id, contact_id,deadline,comment) 
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	`,
		request.Id,
		request.Description,
		fmt.Sprintf("POINT(%f %f)", request.Geo.Lon, request.Geo.Lat),
		request.Address,
		request.CreatedAt,
		request.DeletedAt,
		request.Status,
		request.Priority,
		request.ThemeId,
		request.UserId,
		request.Contact.Id,
		request.Deadline,
		request.Comment,
	)
	if err != nil {
		return nil, errors.Wrap(err, "can not execute create query")
	}
	return request, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*model.Request, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select id, description, ST_AsEWKB(geo), address, created_at, deleted_at, status, priority, theme_id, user_id, contact_id, deadline,comment from requests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.Request
	for rows.Next() {
		var request model.Request
		var contactId uuid.UUID
		var geo []byte
		err := rows.Scan(
			&request.Id,
			&request.Description,
			&geo,
			&request.Address,
			&request.CreatedAt,
			&request.DeletedAt,
			&request.Status,
			&request.Priority,
			&request.ThemeId,
			&request.UserId,
			&contactId,
			&request.Deadline,
			&request.Comment,
		)
		if err != nil {
			return nil, err
		}
		request.Contact = &model.Contact{
			Id: contactId,
		}
		point, err := ewkb.Unmarshal(geo)
		if err != nil {
			return nil, err
		}
		request.Geo = model.GeoPoint{
			Lat: point.FlatCoords()[1],
			Lon: point.FlatCoords()[0],
		}
		requests = append(requests, &request)
	}
	return requests, nil
}

func (r *Repository) GetAllWithThemeIds(ctx context.Context, themeIds []*uuid.UUID) ([]*model.Request, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}

	var themeIdsStrings []string
	for _, themeId := range themeIds {
		themeIdsStrings = append(themeIdsStrings, themeId.String())
	}

	rows, err := connection.Query(
		fmt.Sprintf("select id, description, ST_AsEWKB(geo), address, created_at, deleted_at, status, priority, theme_id, user_id, contact_id, deadline, comment from requests where theme_id in (%s)", fmt.Sprintf("'%s'", strings.Join(themeIdsStrings, "','"))))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*model.Request
	for rows.Next() {
		var request model.Request
		var contactId uuid.UUID
		var geo []byte
		err := rows.Scan(
			&request.Id,
			&request.Description,
			&geo,
			&request.Address,
			&request.CreatedAt,
			&request.DeletedAt,
			&request.Status,
			&request.Priority,
			&request.ThemeId,
			&request.UserId,
			&contactId,
			&request.Deadline,
			&request.Comment,
		)
		if err != nil {
			return nil, err
		}
		request.Contact = &model.Contact{
			Id: contactId,
		}
		point, err := ewkb.Unmarshal(geo)
		if err != nil {
			return nil, err
		}
		request.Geo = model.GeoPoint{
			Lat: point.FlatCoords()[1],
			Lon: point.FlatCoords()[0],
		}
		requests = append(requests, &request)
	}
	return requests, nil
}

func (r *Repository) GetFiles(ctx context.Context, id *uuid.UUID) ([]*uuid.UUID, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select file_id from requests_files where request_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var fileIds []*uuid.UUID
	for rows.Next() {
		var fileId uuid.UUID
		err := rows.Scan(
			&fileId,
		)
		if err != nil {
			return nil, err
		}
		fileIds = append(fileIds, &fileId)
	}
	return fileIds, nil
}

func (r *Repository) AddFile(ctx context.Context, requestId, fileId *uuid.UUID) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		insert into requests_files (request_id, file_id) 
		values ($1,$2)
	`,
		requestId,
		fileId,
	)
	if err != nil {
		return errors.Wrap(err, "can not execute create query")
	}
	return nil
}

func (r *Repository) RemoveFile(ctx context.Context, requestId, fileId *uuid.UUID) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		delete from requests_files where request_id=$1 and file_id=$2
	`,
		requestId,
		fileId,
	)
	if err != nil {
		return errors.Wrap(err, "can not execute delete query")
	}
	return nil
}

func (r *Repository) GetCountWithThemeId(ctx context.Context, from time.Time, to time.Time, themeId *uuid.UUID) (float64, error) {
	connection, err := r.getConnection()
	if err != nil {
		return 0, errors.Wrap(err, "can not get database connection")
	}

	var count float64
	err = connection.QueryRow(`select count(*) from requests where theme_id = $1 and created_at >= $2 and created_at < $3`, themeId, from, to).
		Scan(
			&count,
		)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetCountWithThemeIdAndStatus(ctx context.Context, themeId *uuid.UUID, status int) (float64, error) {
	connection, err := r.getConnection()
	if err != nil {
		return 0, errors.Wrap(err, "can not get database connection")
	}

	var count float64
	err = connection.QueryRow(`select count(*) from requests where theme_id = $1 and status = $2`, themeId, status).
		Scan(
			&count,
		)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) Update(ctx context.Context, request *model.Request) error {
	if request.Id == uuid.Nil {
		return errors.New("request id is nil")
	}
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		update requests set description=$2, geo=$3, address=$4, status=$5, priority=$6, theme_id=$7, user_id=$8, contact_id=$9,deadline=$10,comment=$11
		where id=$1
	`,
		request.Id,
		request.Description,
		fmt.Sprintf("POINT(%f %f)", request.Geo.Lon, request.Geo.Lat),
		request.Address,
		request.Status,
		request.Priority,
		request.ThemeId,
		request.UserId,
		request.Contact.Id,
		request.Deadline,
		request.Comment,
	)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id *uuid.UUID) error {
	if id == nil {
		return errors.New("id is nuk")
	}
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	_, err = connection.Exec(`
		delete from requests where id=$1
	`,
		id,
	)
	if err != nil {
		return errors.Wrap(err, "can not execute delete query")
	}
	return nil
}
