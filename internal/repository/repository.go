package repository

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"time"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context) ([]*model.User, error)
	GetAllWithDepartmentIds(ctx context.Context, departmentIds []string) ([]*model.User, error)
}

type DistrictRepository interface {
	GetByID(ctx context.Context, id string) (*model.District, error)
	Create(ctx context.Context, district *model.District) (*model.District, error)
	GetAll(ctx context.Context) ([]*model.District, error)
}

type WebAuthnRepository interface {
	Set(ctx context.Context, id string, session *webauthn.SessionData) error
	Get(ctx context.Context, id string) (*webauthn.SessionData, error)
	AddUserCredential(ctx context.Context, userId string, credentialId []byte, credential *webauthn.Credential, requestedFrom string) error
	UpdateUserCredential(ctx context.Context, userId string, credentialId []byte, credential *webauthn.Credential) error
	GetUserPublicKeys(ctx context.Context, userId string) ([]model.PublicKey, error)
	GetUserIdByCredentialId(ctx context.Context, id []byte) (string, error)
	DeleteCredentialById(ctx context.Context, id []byte) error
}

type FileRepository interface {
	GetByID(ctx context.Context, id string) (*model.File, error)
	Create(ctx context.Context, file *model.File) (*model.File, error)
	Delete(ctx context.Context, id string) error
}

type DepartmentRepository interface {
	GetByID(ctx context.Context, id string) (*model.Department, error)
	Create(ctx context.Context, department *model.Department) (*model.Department, error)
	GetAll(ctx context.Context) ([]*model.Department, error)
	GetAllWithDistrictId(ctx context.Context, districtId string) ([]*model.Department, error)
	Update(ctx context.Context, department *model.Department) error
}

type ThemeRepository interface {
	GetByID(ctx context.Context, id string) (*model.Theme, error)
	Create(ctx context.Context, theme *model.Theme) (*model.Theme, error)
	GetAll(ctx context.Context, onlyDeleted bool) ([]*model.Theme, error)
	GetAllWithDepartmentIds(ctx context.Context, departmentIds []string, onlyDeleted bool) ([]*model.Theme, error)
	Update(ctx context.Context, theme *model.Theme) error
}

type ContactRepository interface {
	GetByID(ctx context.Context, id string) (*model.Contact, error)
	Create(ctx context.Context, contact *model.Contact) (*model.Contact, error)
	Update(ctx context.Context, contact *model.Contact) error
}

type RequestRepository interface {
	GetByID(ctx context.Context, id string) (*model.Request, error)
	Create(ctx context.Context, request *model.Request) (*model.Request, error)
	GetAll(ctx context.Context) ([]*model.Request, error)
	GetFiles(ctx context.Context, id string) ([]*uuid.UUID, error)
	AddFile(ctx context.Context, requestId, fileId string) error
	RemoveFile(ctx context.Context, requestId, fileId string) error
	GetCountWithThemeId(ctx context.Context, from time.Time, to time.Time, themeId string) (float64, error)
	GetCountWithThemeIdAndStatus(ctx context.Context, themeId string, status int) (float64, error)
	Update(ctx context.Context, request *model.Request) error
}
