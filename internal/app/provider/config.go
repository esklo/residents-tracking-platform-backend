package provider

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/config"
	"go.uber.org/zap"
)

func (s *ServiceProvider) AppConfig() config.AppConfig {
	if s.appConfig == nil {
		cfg, err := config.NewAppConfig()
		if err != nil {
			s.GetLogger().Error("failed to get app config", zap.Error(err))
		}

		s.appConfig = cfg
	}

	return s.appConfig
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			s.GetLogger().Error("failed to get grpc config", zap.Error(err))
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *ServiceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			s.GetLogger().Error("failed to get http config", zap.Error(err))
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *ServiceProvider) DatabaseConfig() config.DatabaseConfig {
	if s.databaseConfig == nil {
		cfg, err := config.NewDatabaseConfig()
		if err != nil {
			s.GetLogger().Error("failed to get database config", zap.Error(err))
		}

		s.databaseConfig = cfg
	}

	return s.databaseConfig
}

func (s *ServiceProvider) S3Config() config.S3Config {
	if s.s3Config == nil {
		cfg, err := config.NewS3Config()
		if err != nil {
			s.GetLogger().Error("failed to get s3 config", zap.Error(err))
		}

		s.s3Config = cfg
	}

	return s.s3Config
}

func (s *ServiceProvider) DadataConfig() config.DadataConfig {
	if s.dadataConfig == nil {
		cfg, err := config.NewDadataConfig()
		if err != nil {
			s.GetLogger().Error("failed to get dadata config", zap.Error(err))
		}

		s.dadataConfig = cfg
	}

	return s.dadataConfig
}
