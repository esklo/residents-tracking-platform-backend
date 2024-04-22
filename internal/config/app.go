package config

import (
	"errors"
	"os"
)

const (
	jwtSecretEnvName     = "JWT_SECRET"
	adminEmailEnvName    = "ADMIN_EMAIL"
	adminPasswordEnvName = "ADMIN_PASSWORD"
	domainEnvName        = "INTERNAL_APP_DOMAIN"
	protocolEnvName      = "INTERNAL_APP_PROTOCOL"
)

type AppConfig interface {
	JwtSecret() string
	AdminEmail() string
	AdminPassword() string
	Domain() string
	Protocol() string
}

type appConfig struct {
	jwtSecret     string
	adminEmail    string
	adminPassword string
	domain        string
	protocol      string
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

	domain := os.Getenv(domainEnvName)
	if len(domain) == 0 {
		return nil, errors.New("domain not found")
	}

	protocol := os.Getenv(protocolEnvName)
	if len(protocol) == 0 {
		return nil, errors.New("protocol not found")
	}

	return &appConfig{
		jwtSecret:     jwtSecret,
		adminEmail:    adminEmail,
		adminPassword: adminPassword,
		domain:        domain,
		protocol:      protocol,
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

func (cfg *appConfig) Domain() string {
	return cfg.domain
}

func (cfg *appConfig) Protocol() string {
	return cfg.protocol
}
