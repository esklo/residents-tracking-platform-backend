package request

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform/gen/proto/request"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/esklo/residents-tracking-platform/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Implementation struct {
	proto.UnimplementedRequestServiceServer
	requestService service.RequestService
	authService    service.AuthService
}

func NewImplementation(requestService service.RequestService, authService service.AuthService) *Implementation {
	return &Implementation{
		requestService: requestService,
		authService:    authService,
	}
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.Request, error) {
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

	request, err := i.requestService.Create(ctx, &themeId, req.Description, req.Address, &model.Contact{
		Phone: req.Contact.Phone,
		Email: req.Contact.Email,
		Name:  req.Contact.Name,
	}, model.GeoPoint{
		Lat: float64(req.Geo.Latitude),
		Lon: float64(req.Geo.Longitude),
	}, fileIds)
	if err != nil {
		return nil, err
	}
	return request.ToProto()
}

func (i Implementation) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	requests, err := i.requestService.GetAll(ctx)
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

func (i Implementation) Update(ctx context.Context, req *proto.Request) (*proto.Request, error) {
	return nil, nil
}

func (i Implementation) GetAsGeoJson(ctx context.Context, req *proto.GetRequest) (*proto.GetAsGeoJsonResponse, error) {
	bytes, err := i.requestService.GetAllAsGeoJson(ctx)
	return &proto.GetAsGeoJsonResponse{Geojson: bytes}, err
}