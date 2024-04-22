package provider

import (
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/auth"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	webauthn2 "github.com/esklo/residents-tracking-platform-backend/internal/repository/webauthn"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	authService "github.com/esklo/residents-tracking-platform-backend/internal/service/auth"
	"github.com/go-webauthn/webauthn/webauthn"
	"log"
	"time"
)

func (s *ServiceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(),
			s.AppConfig(),
			s.WebAuthn(),
			s.WebAuthnRepository(),
		)
	}

	return s.authService
}

func (s *ServiceProvider) AuthImpl() *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService())
	}

	return s.authImpl
}

func (s *ServiceProvider) WebAuthn() *webauthn.WebAuthn {
	if s.webauthn == nil {
		authn, err := webauthn.New(&webauthn.Config{
			Debug:         true,
			RPID:          s.AppConfig().Domain(),
			RPDisplayName: "Система обращений",
			RPOrigins:     []string{fmt.Sprintf("%s://%s", s.AppConfig().Protocol(), s.AppConfig().Domain())},
			Timeouts: webauthn.TimeoutsConfig{
				Login: webauthn.TimeoutConfig{
					Enforce: true,
					Timeout: time.Minute,
				},
				Registration: webauthn.TimeoutConfig{
					Enforce: true,
					Timeout: time.Minute,
				},
			},
		})
		if err != nil {
			log.Fatalf("can not create webauthn: %s", err)
		}
		s.webauthn = authn
	}
	return s.webauthn
}

func (s *ServiceProvider) WebAuthnRepository() repository.WebAuthnRepository {
	if s.webAuthnRepository == nil {
		s.webAuthnRepository = webauthn2.NewRepository(s.DatabaseConnection)
	}

	return s.webAuthnRepository
}
