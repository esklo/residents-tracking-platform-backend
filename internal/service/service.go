package service

import (
	"context"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
	"time"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.User, error)
}

type AuthService interface {
	CreateToken(ctx context.Context, id uuid.UUID) (string, error)
	Login(ctx context.Context, email, password string) (*model.User, error)
	ChangePassword(ctx context.Context, user *model.User, password string) error

	RequestPublicKeyAttestation(ctx context.Context, user *model.User) (flowId, credential []byte, err error)
	PublicKeyAttestation(ctx context.Context, user *model.User, flowId []byte, credential *protocol.CredentialCreationResponse) ([]byte, error)
	RequestPublicKeyAssertion(ctx context.Context) ([]byte, []byte, error)
	PublicKeyAssertion(ctx context.Context, flowId []byte, credential *protocol.CredentialAssertionResponse) (*model.User, error)

	GetPublicKeys(ctx context.Context, user *model.User) ([]model.PublicKey, error)
	DeletePublicKey(ctx context.Context, id []byte) error

	ExchangeToken(ctx context.Context, tokenString string) (*model.User, error)
	ExchangeTokenFromContext(ctx context.Context) (*model.User, error)
}

type GeoService interface {
	BuildingsByCoordinates(ctx context.Context, lat, lon float64) ([]*model.GeoBuilding, error)
	GetDistricts(ctx context.Context, areaId int64, level int) ([]*model.GeoDistrict, error)
	Suggest(ctx context.Context, query string) ([]*suggest.AddressSuggestion, error)
	Locate(ctx context.Context, lat, lon float64) (*suggest.AddressSuggestion, error)
}

type DistrictService interface {
	Create(ctx context.Context, areaId int64, fileId *uuid.UUID) (*model.District, error)
	Get(ctx context.Context, id *uuid.UUID) (*model.District, error)
	GetAll(ctx context.Context) ([]*model.District, error)
}

type FileService interface {
	Upload(ctx context.Context, filename string, data []byte) (*model.File, error)
	GetById(ctx context.Context, id *uuid.UUID) (*model.File, error)
	Delete(ctx context.Context, id *uuid.UUID) error
}

type DepartmentService interface {
	Create(ctx context.Context, department *model.Department) (*model.Department, error)
	Get(ctx context.Context, id *uuid.UUID) (*model.Department, error)
	GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.Department, error)
	Update(ctx context.Context, department *model.Department) error
}

type ThemeService interface {
	Create(ctx context.Context, department *model.Theme) (*model.Theme, error)
	Get(ctx context.Context, id *uuid.UUID) (*model.Theme, error)
	GetAll(ctx context.Context, districtId *uuid.UUID) ([]*model.Theme, error)
	Update(ctx context.Context, theme *model.Theme) error
	Delete(ctx context.Context, id *uuid.UUID) error
	GetAllWithDepartment(ctx context.Context, department *uuid.UUID) ([]*model.Theme, error)
}

type ContactService interface {
	Create(ctx context.Context, contact *model.Contact) (*model.Contact, error)
	Update(ctx context.Context, contact *model.Contact) error
	Get(ctx context.Context, id *uuid.UUID) (*model.Contact, error)
}

type RequestService interface {
	Create(ctx context.Context, themeId *uuid.UUID, description, address string, contact *model.Contact, geo model.GeoPoint, fileIds []*uuid.UUID, deadline *time.Time) (*model.Request, error)
	Get(ctx context.Context, id *uuid.UUID) (*model.Request, error)
	GetAll(ctx context.Context) ([]*model.Request, error)
	GetAllAsGeoJson(ctx context.Context) ([]byte, error)
	GetCountWithThemeId(ctx context.Context, from time.Time, to time.Time, themeId string) (float64, error)
	GetCountWithThemeIdAndStatus(ctx context.Context, themeId string, status model.RequestStatus) (float64, error)
	Update(ctx context.Context, request *model.Request, fileIds []*uuid.UUID) error
	ExportExcel(ctx context.Context) (*model.File, error)
}

type AnalyticsService interface {
	RequestsPerTheme(ctx context.Context, from time.Time, to time.Time, departmentId *uuid.UUID) ([]*model.RequestsPerTheme, error)
	Stats(ctx context.Context, departmentId *uuid.UUID) (*model.StatsElement, error)
	RequestsPerThemePerDate(ctx context.Context, departmentId *uuid.UUID, from time.Time, to time.Time) ([]*model.RequestsPerThemePerDateElement, error)
}
