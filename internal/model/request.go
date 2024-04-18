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
