package geo

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"go.uber.org/zap"
)

var _ def.GeoService = (*Service)(nil)

type Service struct {
	dadataClient *suggest.Api
	logger       *zap.Logger
}

func NewService(dadataClient *suggest.Api, logger *zap.Logger) *Service {
	return &Service{
		dadataClient: dadataClient,
		logger:       logger,
	}
}
