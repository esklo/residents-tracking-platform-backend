package request

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/request"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Implementation struct {
	proto.UnimplementedRequestServiceServer
	requestService service.RequestService
	authService    service.AuthService
	logger         *zap.Logger
}

func NewImplementation(requestService service.RequestService, authService service.AuthService, logger *zap.Logger) *Implementation {
	return &Implementation{
		requestService: requestService,
		authService:    authService,
		logger:         logger,
	}
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Request, error) {
	i.logger.Debug("request.GetById request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	requestId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse request id")
	}
	request, err := i.requestService.Get(ctx, &requestId)
	if err != nil {
		return nil, err
	}
	return request.ToProto()
}

func (i Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.Request, error) {
	i.logger.Debug("request.Create request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	themeId, err := uuid.Parse(req.ThemeId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse theme id")
	}

	var fileIds []*uuid.UUID
	for _, file := range req.Files {
		fileId, err := uuid.Parse(file)
		if err != nil {
			return nil, errors.Wrap(err, "can not parse file id")
		}
		fileIds = append(fileIds, &fileId)
	}

	var deadline *time.Time
	if req.Deadline != nil {
		d := req.Deadline.AsTime()
		deadline = &d
	}

	request, err := i.requestService.Create(ctx, &themeId, req.Description, req.Address, &model.Contact{
		Phone: req.Contact.Phone,
		Email: req.Contact.Email,
		Name:  req.Contact.Name,
	}, model.GeoPoint{
		Lat: float64(req.Geo.Latitude),
		Lon: float64(req.Geo.Longitude),
	}, fileIds, deadline)
	if err != nil {
		return nil, err
	}
	return request.ToProto()
}

func (i Implementation) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	i.logger.Debug("request.Get request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	requests, err := i.requestService.GetAllWithDepartment(ctx, user.DepartmentId)
	if err != nil {
		return nil, err
	}
	var protoRequests []*proto.Request
	for _, request := range requests {
		protoRequest, err := request.ToProto()
		if err != nil {
			return nil, err
		}
		protoRequests = append(protoRequests, protoRequest)
	}
	return &proto.GetResponse{Requests: protoRequests}, nil
}

func (i Implementation) Update(ctx context.Context, req *proto.UpdateRequest) (*empty.Empty, error) {
	i.logger.Debug("request.Update request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	requestId, err := uuid.Parse(req.Request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse request id")
	}

	themeId, err := uuid.Parse(req.Request.ThemeId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse theme id")
	}

	var userId *uuid.UUID
	if req.Request.UserId != "" {
		userIdP, err := uuid.Parse(req.Request.UserId)
		if err != nil {
			return nil, errors.Wrap(err, "can not parse user id")
		}
		userId = &userIdP
	}

	var contactId uuid.UUID
	if req.Request.Contact.Id != "" {
		contactId, err = uuid.Parse(req.Request.Contact.Id)
		if err != nil {
			return nil, errors.Wrap(err, "can not parse contact id")
		}
	}

	var fileIds []*uuid.UUID
	for _, file := range req.Files {
		fileId, err := uuid.Parse(file)
		if err != nil {
			return nil, errors.Wrap(err, "can not parse file id")
		}
		fileIds = append(fileIds, &fileId)
	}

	var requestStatus = model.RequestStatusUnknown
	switch req.Request.Status {
	case 1:
		requestStatus = model.RequestStatusOpen
		break
	case 2:
		requestStatus = model.RequestStatusClosed
		break
	case 3:
		requestStatus = model.RequestStatusDeclined
		break
	}

	var requestPriority = model.RequestPriorityUnknown
	switch req.Request.Priority {
	case 1:
		requestPriority = model.RequestPriorityDefault
		break
	case 2:
		requestPriority = model.RequestPriorityLow
		break
	case 3:
		requestPriority = model.RequestPriorityHigh
		break
	}

	var deadline *time.Time
	if req.Request.Deadline != nil {
		d := req.Request.Deadline.AsTime()
		deadline = &d
	}

	err = i.requestService.Update(ctx, &model.Request{
		Id:          requestId,
		Description: req.Request.Description,
		Geo: model.GeoPoint{
			Lat: float64(req.Request.Geo.Latitude),
			Lon: float64(req.Request.Geo.Longitude),
		},
		Address:  req.Request.Address,
		Status:   requestStatus,
		Priority: requestPriority,
		ThemeId:  &themeId,
		Contact: &model.Contact{
			Id:    contactId,
			Phone: req.Request.Contact.Phone,
			Email: req.Request.Contact.Email,
			Name:  req.Request.Contact.Name,
		},
		UserId:   userId,
		Deadline: deadline,
	}, fileIds)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (i Implementation) GetAsGeoJson(ctx context.Context, req *proto.GetRequest) (*proto.GetAsGeoJsonResponse, error) {
	i.logger.Debug("request.GetAsGeoJson request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	bytes, err := i.requestService.GetAllAsGeoJsonWithDepartment(ctx, user.DepartmentId)
	return &proto.GetAsGeoJsonResponse{Geojson: bytes}, err
}

func (i Implementation) ExportExcel(ctx context.Context, req *empty.Empty) (*proto.ExportResponse, error) {
	i.logger.Debug("request.ExportExcel request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	file, err := i.requestService.ExportExcel(ctx, user.DepartmentId)
	if err != nil {
		return nil, errors.Wrap(err, "can not export excel")
	}
	protoFile, err := file.ToProto()
	if err != nil {
		return nil, errors.Wrap(err, "can not convert file to proto")
	}
	return &proto.ExportResponse{File: protoFile}, nil
}
