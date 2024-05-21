package district

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/district"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Implementation struct {
	proto.UnimplementedDistrictServiceServer
	districtService service.DistrictService
	authService     service.AuthService
	logger          *zap.Logger
}

func NewImplementation(districtService service.DistrictService, authService service.AuthService, logger *zap.Logger) *Implementation {
	return &Implementation{
		districtService: districtService,
		authService:     authService,
		logger:          logger,
	}
}

func (i Implementation) Get(ctx context.Context, _ *proto.GetRequest) (*proto.GetResponse, error) {
	i.logger.Debug("district.Get request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	if !user.IsAdmin() {
		return nil, model.ErrorPermissionDenied
	}

	districts, err := i.districtService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var protoDistricts []*proto.District
	for _, district := range districts {
		protoDistrict, err := district.ToProto()
		if err != nil {
			return nil, err
		}
		protoDistricts = append(protoDistricts, protoDistrict)
	}
	return &proto.GetResponse{
		Districts: protoDistricts,
	}, nil
}

func (i Implementation) GetById(ctx context.Context, req *proto.ByIdRequest) (*proto.District, error) {
	i.logger.Debug("district.GetById request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	districtId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse coat of arms file id")
	}
	district, err := i.districtService.Get(ctx, &districtId)
	if err != nil {
		return nil, err
	}
	return district.ToProto()
}

func (i Implementation) Create(ctx context.Context, req *proto.CreateRequest) (*proto.District, error) {
	i.logger.Debug("district.Create request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	if !user.IsAdmin() {
		return nil, model.ErrorPermissionDenied
	}

	var coatOfArmsFileId *uuid.UUID
	if req.CoatOfArmsFileId != nil {
		coatOfArmsFileIdParsed, err := uuid.Parse(*req.CoatOfArmsFileId)
		if err != nil {
			return nil, errors.Wrap(err, "can not parse coat of arms file id")
		}
		coatOfArmsFileId = &coatOfArmsFileIdParsed
	}
	if req.AreaId == 0 {
		return nil, errors.New("areaId is required")
	}
	district, err := i.districtService.Create(ctx, req.AreaId, coatOfArmsFileId)
	if err != nil {
		return nil, err
	}

	protoDistrict, err := district.ToProto()
	if err != nil {
		return nil, err
	}
	return protoDistrict, nil
}
