package model

import (
	protoAuth "github.com/esklo/residents-tracking-platform/gen/proto/auth"
	"github.com/go-webauthn/webauthn/webauthn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PublicKey struct {
	Credential    webauthn.Credential
	CredentialId  []byte
	CreatedAt     *time.Time
	LastUsedAt    *time.Time
	RequestedFrom string
}

func (p *PublicKey) ToProto() (*protoAuth.Key, error) {
	if p == nil {
		return nil, ErrorModelIsEmpty
	}
	key := protoAuth.Key{
		Id: p.CredentialId,
	}
	if p.CreatedAt != nil {
		key.CreatedAt = timestamppb.New(*p.CreatedAt)
	}
	if p.LastUsedAt != nil {
		key.LastUsedAt = timestamppb.New(*p.LastUsedAt)
	}
	if p.RequestedFrom != "" {
		key.RequestedFrom = &p.RequestedFrom
	}
	return &key, nil
}
