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
		return nil, errors.Wrap(err, "can not get request by id")
	}
	return request.ToProto()
}

func (i Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.Request, error) {
	i.logger.Debug("request.Create request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	req.Request.Id = uuid.Nil.String()

	requestM := model.Request{}
	if err := requestM.FromProto(req.Request); err != nil {
		return nil, err
	}

	request, err := i.requestService.Create(ctx, &requestM)
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

	var request model.Request
	if err := request.FromProto(req.Request); err != nil {
		return nil, err
	}

	err = i.requestService.Update(ctx, &request)
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

func (i Implementation) Delete(ctx context.Context, req *proto.ByIdRequest) (*empty.Empty, error) {
	i.logger.Debug("request.Delete request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	requestId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse request id")
	}

	if err := i.requestService.Delete(ctx, &requestId); err != nil {
		return nil, errors.Wrap(err, "can not delete request")
	}

	return nil, nil
}
