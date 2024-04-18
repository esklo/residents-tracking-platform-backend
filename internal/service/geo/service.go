package geo

import (
	"github.com/ekomobile/dadata/v2/api/suggest"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
)

var _ def.GeoService = (*Service)(nil)

type Service struct {
	dadataClient *suggest.Api
}

func NewService(dadataClient *suggest.Api) *Service {
	return &Service{
		dadataClient: dadataClient,
	}
}
