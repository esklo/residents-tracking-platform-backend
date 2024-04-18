package auth

import (
	"github.com/esklo/residents-tracking-platform-backend/internal/config"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/go-webauthn/webauthn/webauthn"
)

var _ def.AuthService = (*Service)(nil)

type Service struct {
	userRepository     repository.UserRepository
	appConfig          config.AppConfig
	webAuthn           *webauthn.WebAuthn
	webAuthnRepository repository.WebAuthnRepository
}

func NewService(
	userRepository repository.UserRepository,
	appConfig config.AppConfig,
	webAuthn *webauthn.WebAuthn,
	webAuthnRepository repository.WebAuthnRepository,
) *Service {
	return &Service{
		userRepository:     userRepository,
		appConfig:          appConfig,
		webAuthn:           webAuthn,
		webAuthnRepository: webAuthnRepository,
	}
}
