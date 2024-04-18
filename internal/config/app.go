package config

import (
	"errors"
	"os"
)

const (
	jwtSecretEnvName     = "JWT_SECRET"
	adminEmailEnvName    = "ADMIN_EMAIL"
	adminPasswordEnvName = "ADMIN_PASSWORD"
)

type AppConfig interface {
	JwtSecret() string
	AdminEmail() string
	AdminPassword() string
}

type appConfig struct {
	jwtSecret     string
	adminEmail    string
	adminPassword string
}

func NewAppConfig() (AppConfig, error) {
	jwtSecret := os.Getenv(jwtSecretEnvName)
	if len(jwtSecret) == 0 {
		return nil, errors.New("jwt secret not found")
	}

	adminEmail := os.Getenv(adminEmailEnvName)
	if len(adminEmail) == 0 {
		return nil, errors.New("admin email not found")
	}

	adminPassword := os.Getenv(adminPasswordEnvName)
	if len(adminPassword) == 0 {
		return nil, errors.New("admin password not found")
	}

	return &appConfig{
		jwtSecret:     jwtSecret,
		adminEmail:    adminEmail,
		adminPassword: adminPassword,
	}, nil
}

func (cfg *appConfig) JwtSecret() string {
	return cfg.jwtSecret
}

func (cfg *appConfig) AdminEmail() string {
	return cfg.adminEmail
}

func (cfg *appConfig) AdminPassword() string {
	return cfg.adminPassword
}
