package geo

import (
	"context"
	"github.com/ekomobile/dadata/v2/api/suggest"
)

func (s *Service) Suggest(ctx context.Context, query string) ([]*suggest.AddressSuggestion, error) {
	suggestions, err := s.dadataClient.Address(ctx, &suggest.RequestParams{
		Query: query,
		Count: 5,
	})
	if err != nil {
		return nil, err
	}
	return suggestions, nil
}
