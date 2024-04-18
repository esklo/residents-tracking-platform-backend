package model

import (
	"crypto/rand"
	"encoding/hex"
	protoUser "github.com/esklo/residents-tracking-platform-backend/gen/proto/user"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Id           uuid.UUID
	Email        string
	FirstName    string
	LastName     *string
	FatherName   *string
	Password     string
	Salt         string
	Role         UserRole
	DepartmentId *uuid.UUID

	webAuthnCredentials []webauthn.Credential
	webAuthnID          []byte
	//todo: department
}

func (u *User) SetPassword(password string) error {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}
	saltedPassword := append(salt, []byte(password)...)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, 14)
	if err != nil {
		return err
	}
	u.Password = hex.EncodeToString(hashedPassword)
	u.Salt = hex.EncodeToString(salt)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	salt, err := hex.DecodeString(u.Salt)
	if err != nil {
		return false
	}
	hashedPassword, err := hex.DecodeString(u.Password)
	if err != nil {
		return false
	}
	saltedPassword := append(salt, []byte(password)...)
	err = bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
	return err == nil
}

func (u *User) ToProto() (*protoUser.User, error) {
	if u == nil {
		return nil, ErrorModelIsEmpty
	}
	userProto := protoUser.User{
		Id:         u.Id.String(),
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		FatherName: u.FatherName,
	}
	switch u.Role {
	case UserRoleAdmin:
		userProto.Role = protoUser.Role_AdminRole
	default:
		userProto.Role = protoUser.Role_EmployeeRole
	}
	if u.DepartmentId != nil {
		departmentId := u.DepartmentId.String()
		userProto.DepartmentId = &departmentId
	}
	return &userProto, nil
}

func (u *User) WebAuthnID() []byte {
	return u.webAuthnID[:]
}

func (u *User) WebAuthnName() string {
	return u.Email
}

func (u *User) WebAuthnDisplayName() string {
	var nameParts []string
	if u.LastName != nil {
		nameParts = append(nameParts, *u.LastName)
	}
	nameParts = append(nameParts, u.FirstName)
	if u.FatherName != nil {
		nameParts = append(nameParts, *u.FatherName)
	}
	return strings.Join(nameParts, " ")
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.webAuthnCredentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) SetWebAuthnCredentials(credentials []webauthn.Credential) {
	u.webAuthnCredentials = credentials
}

func (u *User) SetWebAuthnID(id []byte) {
	u.webAuthnID = id
}

func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}
