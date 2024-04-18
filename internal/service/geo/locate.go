package geo

import (
	"context"
	"fmt"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/pkg/errors"
	"net/url"
)

func (s *Service) Locate(ctx context.Context, lat, lon float64) (*suggest.AddressSuggestion, error) {
	var result = &suggest.AddressResponse{}
	err := s.dadataClient.Client.Get(ctx, "geolocate/address", url.Values{
		"lat":   {fmt.Sprintf("%f", lat)},
		"lon":   {fmt.Sprintf("%f", lon)},
		"count": {"1"},
	}, result)
	if err != nil {
		return nil, err
	}
	if len(result.Suggestions) == 0 {
		return nil, errors.New("not found")
	}
	return result.Suggestions[0], nil
}
