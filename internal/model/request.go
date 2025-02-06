package model

import (
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/file"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/geo"
	protoRequest "github.com/esklo/residents-tracking-platform-backend/gen/proto/request"
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Request struct {
	Id          uuid.UUID       `json:"id,omitempty"`
	Description string          `json:"description,omitempty"`
	Geo         GeoPoint        `json:"geo"`
	Address     string          `json:"address,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
	DeletedAt   *time.Time      `json:"deletedAt,omitempty"`
	Status      RequestStatus   `json:"status,omitempty"`
	Priority    RequestPriority `json:"priority,omitempty"`
	ThemeId     *uuid.UUID      `json:"themeId,omitempty"`
	Contact     *Contact        `json:"contact,omitempty"`
	UserId      *uuid.UUID      `json:"userId,omitempty"`
	Files       []*File         `json:"files,omitempty"`
	ReportFiles []*File         `json:"report_files,omitempty"`
	Deadline    *time.Time      `json:"deadline,omitempty"`
	Comment     *string         `json:"comment,omitempty"`
	//todo
}

func (r *Request) ToProto() (*protoRequest.Request, error) {
	if r == nil {
		return nil, ErrorModelIsEmpty
	}
	request := protoRequest.Request{
		Id:          r.Id.String(),
		Description: r.Description,
		Geo: &geo.GeoPoint{
			Latitude:  float32(r.Geo.Lat),
			Longitude: float32(r.Geo.Lon),
		},
		Address:   r.Address,
		CreatedAt: timestamppb.New(r.CreatedAt),
		Comment:   r.Comment,
	}
	if r.ThemeId != nil {
		request.ThemeId = r.ThemeId.String()
	}
	if r.Contact != nil {
		protoContact, err := r.Contact.ToProto()
		if err != nil {
			return nil, err
		}
		request.Contact = protoContact
	}
	if r.UserId != nil {
		request.UserId = r.UserId.String()
	}
	if r.DeletedAt != nil {
		request.DeletedAt = timestamppb.New(*r.DeletedAt)
	}
	if r.Deadline != nil {
		request.Deadline = timestamppb.New(*r.Deadline)
	}
	switch r.Status {
	case RequestStatusOpen:
		request.Status = protoRequest.Status_StatusOpen
	case RequestStatusClosed:
		request.Status = protoRequest.Status_StatusClosed
	case RequestStatusDeclined:
		request.Status = protoRequest.Status_StatusDeclined
	default:
		request.Status = protoRequest.Status_StatusUnknown
	}
	switch r.Priority {
	case RequestPriorityDefault:
		request.Priority = protoRequest.Priority_PriorityDefault
	case RequestPriorityHigh:
		request.Priority = protoRequest.Priority_PriorityHigh
	case RequestPriorityLow:
		request.Priority = protoRequest.Priority_PriorityLow
	default:
		request.Priority = protoRequest.Priority_PriorityUnknown
	}
	var protoFiles []*proto.File
	for _, file := range r.Files {
		protoFile, err := file.ToProto()
		if err != nil {
			return nil, err
		}
		protoFiles = append(protoFiles, protoFile)
	}
	request.Files = protoFiles

	var protoReportFiles []*proto.File
	for _, file := range r.ReportFiles {
		protoReportFile, err := file.ToProto()
		if err != nil {
			return nil, err
		}
		protoReportFiles = append(protoReportFiles, protoReportFile)
	}
	request.ReportFiles = protoReportFiles
	return &request, nil
}

func (r *Request) AsGeoJson() (*geojson.Feature, error) {
	if r == nil {
		return nil, errors.New("request is nil")
	}
	feature := geojson.NewFeature(orb.Point{r.Geo.Lon, r.Geo.Lat})
	feature.ID = r.Id.String()
	feature.Properties = geojson.Properties{
		"request": r,
	}
	return feature, nil
}

func (r *Request) FromProto(req *protoRequest.Request) error {
	requestId, err := uuid.Parse(req.Id)
	if err != nil {
		return errors.Wrap(err, "can not parse request id")
	}
	r.Id = requestId

	r.Description = req.Description
	r.Geo = GeoPoint{
		Lat: float64(req.Geo.Latitude),
		Lon: float64(req.Geo.Longitude),
	}
	r.Address = req.Address
	r.CreatedAt = req.CreatedAt.AsTime()

	if req.DeletedAt != nil {
		deletedAt := req.DeletedAt.AsTime()
		r.DeletedAt = &deletedAt
	}
	switch req.Status {
	case protoRequest.Status_StatusOpen:
		r.Status = RequestStatusOpen
	case protoRequest.Status_StatusClosed:
		r.Status = RequestStatusClosed
	case protoRequest.Status_StatusDeclined:
		r.Status = RequestStatusDeclined
	default:
		r.Status = RequestStatusUnknown
	}
	switch req.Priority {
	case protoRequest.Priority_PriorityDefault:
		r.Priority = RequestPriorityDefault
	case protoRequest.Priority_PriorityLow:
		r.Priority = RequestPriorityLow
	case protoRequest.Priority_PriorityHigh:
		r.Priority = RequestPriorityHigh
	default:
		r.Priority = RequestPriorityUnknown
	}
	themeId, err := uuid.Parse(req.ThemeId)
	if err != nil {
		return errors.Wrap(err, "can not parse theme id")
	}
	r.ThemeId = &themeId

	var contact Contact
	if err := contact.FromProto(req.Contact); err != nil {
		return errors.Wrap(err, "can not convert contact from proto")
	}
	r.Contact = &contact
	if userId, err := uuid.Parse(req.UserId); err == nil {
		r.UserId = &userId
	}
	if req.Deadline != nil {
		deadline := req.Deadline.AsTime()
		r.Deadline = &deadline
	}

	r.Comment = req.Comment

	var files []*File
	for _, file := range req.GetFiles() {
		var fileC File
		if err := fileC.FromProto(file); err != nil {
			return errors.Wrap(err, "can not convert file from proto")
		}
		files = append(files, &fileC)
	}
	r.Files = files

	var reportFiles []*File
	for _, file := range req.GetReportFiles() {
		var fileC File
		if err := fileC.FromProto(file); err != nil {
			return errors.Wrap(err, "can not convert file from proto")
		}
		reportFiles = append(reportFiles, &fileC)
	}
	r.ReportFiles = reportFiles
	return nil
}
