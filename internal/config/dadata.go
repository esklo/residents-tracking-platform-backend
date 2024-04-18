package config

import (
	"errors"
	"os"
)

const (
	dadataApiKeyEnvName    = "DADATA_API_KEY"
	dadataSecretKeyEnvName = "DADATA_SECRET_KEY"
)

type DadataConfig interface {
	ApiKey() string
	SecretKey() string
}

type dadataConfig struct {
	apiKey    string
	secretKey string
}

func NewDadataConfig() (DadataConfig, error) {
	apiKey := os.Getenv(dadataApiKeyEnvName)
	if len(apiKey) == 0 {
		return nil, errors.New("dadata api key not found")
	}

	secretKey := os.Getenv(dadataSecretKeyEnvName)
	if len(secretKey) == 0 {
		return nil, errors.New("dadata secret key not found")
	}

	return &dadataConfig{
		apiKey:    apiKey,
		secretKey: secretKey,
	}, nil
}

func (cfg *dadataConfig) ApiKey() string {
	return cfg.apiKey
}

func (cfg *dadataConfig) SecretKey() string {
	return cfg.secretKey
}
