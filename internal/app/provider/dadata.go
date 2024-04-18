package provider

import (
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

func (s *ServiceProvider) DadataClient() *suggest.Api {
	if s.dadataClient == nil {
		dadataClient := dadata.NewSuggestApi(client.WithCredentialProvider(&client.Credentials{
			ApiKeyValue:    s.DadataConfig().ApiKey(),
			SecretKeyValue: s.DadataConfig().SecretKey(),
		}))
		s.dadataClient = dadataClient
	}
	return s.dadataClient
}
