package webauthn

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	def "github.com/esklo/residents-tracking-platform-backend/internal/repository"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/pkg/errors"
	"sync"
	"time"
)

var _ def.WebAuthnRepository = (*Repository)(nil)

type Repository struct {
	mx            sync.RWMutex
	data          map[string]*webauthn.SessionData
	getConnection func() (*sql.DB, error)
}

func NewRepository(getConnection func() (*sql.DB, error)) *Repository {
	return &Repository{
		data:          make(map[string]*webauthn.SessionData),
		getConnection: getConnection,
	}
}

func (r *Repository) cleanup() {
	r.mx.Lock()
	defer r.mx.Unlock()

	for k, v := range r.data {
		if v.Expires.Before(time.Now()) {
			delete(r.data, k)
		}
	}
}

func (r *Repository) Set(_ context.Context, id string, session *webauthn.SessionData) error {
	r.cleanup()

	r.mx.Lock()
	defer r.mx.Unlock()

	r.data[id] = session
	return nil
}

func (r *Repository) Get(_ context.Context, id string) (*webauthn.SessionData, error) {
	r.cleanup()

	r.mx.Lock()
	defer r.mx.Unlock()

	if data, ok := r.data[id]; ok {
		return data, nil
	}
	return nil, errors.New("session data not found")
}

func (r *Repository) AddUserCredential(_ context.Context, userId string, credentialId []byte, credential *webauthn.Credential, requestedFrom string) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	credentialBytes, err := json.Marshal(credential)
	if err != nil {
		return errors.Wrap(err, "can not marshal credential")
	}

	_, err = connection.Exec(`
		insert into user_credentials (user_id, credential_id, credential,created_at,requested_from) 
		values ($1,$2,$3,$4,$5)
	`, userId, credentialId, credentialBytes, time.Now(), requestedFrom)
	if err != nil {
		return errors.Wrap(err, "can not execute create query")
	}
	return nil
}

func (r *Repository) GetUserPublicKeys(ctx context.Context, userId string) ([]model.PublicKey, error) {
	connection, err := r.getConnection()
	if err != nil {
		return nil, errors.Wrap(err, "can not get database connection")
	}
	rows, err := connection.Query("select credential,credential_id,created_at,last_used_at,requested_from from user_credentials where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []model.PublicKey
	for rows.Next() {
		var credentialBytes []byte
		var credentialId []byte
		var createdAt *time.Time
		var lastUsedAt *time.Time
		var requestedFrom string
		err := rows.Scan(&credentialBytes, &credentialId, &createdAt, &lastUsedAt, &requestedFrom)
		if err != nil {
			return nil, err
		}
		var credential webauthn.Credential
		if err := json.Unmarshal(credentialBytes, &credential); err != nil {
			return nil, err
		}
		credentials = append(credentials, model.PublicKey{
			CredentialId:  credentialId,
			Credential:    credential,
			CreatedAt:     createdAt,
			LastUsedAt:    lastUsedAt,
			RequestedFrom: requestedFrom,
		})
	}
	return credentials, nil
}

func (r *Repository) UpdateUserCredential(_ context.Context, userId string, credentialId []byte, credential *webauthn.Credential) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}

	credentialBytes, err := json.Marshal(credential)
	if err != nil {
		return errors.Wrap(err, "can not marshal credential")
	}

	_, err = connection.Exec(`
		update user_credentials set credential = $1, last_used_at = $2 where user_id = $3 and credential_id = $4
	`, credentialBytes, time.Now(), userId, credentialId)
	if err != nil {
		return errors.Wrap(err, "can not execute update query")
	}
	return nil
}

func (r *Repository) GetUserIdByCredentialId(_ context.Context, id []byte) (string, error) {
	connection, err := r.getConnection()
	if err != nil {
		return "", errors.Wrap(err, "can not get database connection")
	}
	var userId string
	err = connection.QueryRow(`select user_id from user_credentials where credential_id = $1`, id).Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (r *Repository) DeleteCredentialById(_ context.Context, id []byte) error {
	connection, err := r.getConnection()
	if err != nil {
		return errors.Wrap(err, "can not get database connection")
	}
	_, err = connection.Exec(`delete from user_credentials where credential_id = $1`, id)
	if err != nil {
		return errors.Wrap(err, "can not delete credential by id")
	}
	return nil
}
