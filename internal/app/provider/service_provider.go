package provider

import (
	"database/sql"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/auth"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/contact"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/department"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/district"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/file"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/geo"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/request"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/theme"
	"github.com/esklo/residents-tracking-platform-backend/internal/api/user"
	"github.com/esklo/residents-tracking-platform-backend/internal/config"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	"github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/esklo/residents-tracking-platform-backend/internal/storage/s3"
	"github.com/go-webauthn/webauthn/webauthn"
)

type ServiceProvider struct {
	appConfig      config.AppConfig
	grpcConfig     config.GRPCConfig
	httpConfig     config.HTTPConfig
	databaseConfig config.DatabaseConfig
	s3Config       config.S3Config
	dadataConfig   config.DadataConfig

	userRepository       repository.UserRepository
	districtRepository   repository.DistrictRepository
	webAuthnRepository   repository.WebAuthnRepository
	fileRepository       repository.FileRepository
	departmentRepository repository.DepartmentRepository
	themeRepository      repository.ThemeRepository
	requestRepository    repository.RequestRepository
	contactRepository    repository.ContactRepository

	userService       service.UserService
	authService       service.AuthService
	geoService        service.GeoService
	districtService   service.DistrictService
	fileService       service.FileService
	departmentService service.DepartmentService
	themeService      service.ThemeService
	requestService    service.RequestService
	contactService    service.ContactService

	userImpl       *user.Implementation
	authImpl       *auth.Implementation
	geoImpl        *geo.Implementation
	districtImpl   *district.Implementation
	fileImpl       *file.Implementation
	departmentImpl *department.Implementation
	themeImpl      *theme.Implementation
	requestImpl    *request.Implementation
	contactImpl    *contact.Implementation

	databaseConnection *sql.DB
	webauthn           *webauthn.WebAuthn

	fileStorage  *s3.Storage
	dadataClient *suggest.Api
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
