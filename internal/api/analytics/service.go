package analytics

import (
	"context"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/analytics"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"go.uber.org/zap"
)

type Implementation struct {
	proto.UnimplementedAnalyticsServiceServer
	authService      service.AuthService
	analyticsService service.AnalyticsService
	logger           *zap.Logger
}

func NewImplementation(authService service.AuthService, analyticsService service.AnalyticsService, logger *zap.Logger) *Implementation {
	return &Implementation{
		authService:      authService,
		analyticsService: analyticsService,
		logger:           logger,
	}
}

func (i Implementation) RequestsPerTheme(ctx context.Context, req *proto.RequestsPerThemeRequest) (*proto.RequestsPerThemeResponse, error) {
	i.logger.Debug("analytics.RequestsPerTheme request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}
	requestsPerTheme, err := i.analyticsService.RequestsPerTheme(ctx, req.From.AsTime(), req.To.AsTime(), user.DepartmentId)
	if err != nil {
		return nil, err
	}
	var requestsPerThemeProto []*proto.RequestsPerTheme
	for _, perTheme := range requestsPerTheme {
		perThemeProto, err := perTheme.ToProto()
		if err != nil {
			return nil, err
		}
		requestsPerThemeProto = append(requestsPerThemeProto, perThemeProto)
	}
	return &proto.RequestsPerThemeResponse{Data: requestsPerThemeProto}, nil
}

func (i Implementation) Stats(ctx context.Context, _ *empty.Empty) (*proto.StatsResponse, error) {
	i.logger.Debug("analytics.Stats request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	stats, err := i.analyticsService.Stats(ctx, user.DepartmentId)
	if err != nil {
		return nil, err
	}
	statsProto, err := stats.ToProto()
	if err != nil {
		return nil, err
	}
	return &proto.StatsResponse{Data: statsProto}, nil
}

func (i Implementation) RequestsPerThemePerDate(ctx context.Context, req *proto.RequestsPerThemeRequest) (*proto.RequestsPerThemePerDateResponse, error) {
	i.logger.Debug("analytics.RequestsPerThemePerDate request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}
	data, err := i.analyticsService.RequestsPerThemePerDate(ctx, user.DepartmentId, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		return nil, err
	}
	var dataProto []*proto.RequestsPerThemePerDate
	for _, dataElement := range data {
		dataElementProto, err := dataElement.ToProto()
		if err != nil {
			return nil, err
		}
		dataProto = append(dataProto, dataElementProto)
	}
	return &proto.RequestsPerThemePerDateResponse{Data: dataProto}, nil
}
