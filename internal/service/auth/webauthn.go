package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/mileusna/useragent"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"log"
)

func (s *Service) RequestPublicKeyAttestation(ctx context.Context, user *model.User) ([]byte, []byte, error) {
	authSelect := protocol.AuthenticatorSelection{
		RequireResidentKey: protocol.ResidentKeyRequired(),
		ResidentKey:        protocol.ResidentKeyRequirementRequired,
	}

	flowId := uuid.New()
	user.SetWebAuthnID(flowId[:])

	creation, data, err := s.webAuthn.BeginRegistration(user, webauthn.WithAuthenticatorSelection(authSelect))
	if err != nil {
		return nil, nil, errors.Wrap(err, "can not create challenge")
	}

	if err := s.webAuthnRepository.Set(ctx, flowId.String(), data); err != nil {
		return nil, nil, errors.Wrap(err, "can not save session")
	}

	credentialBytes, err := json.Marshal(creation.Response)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can not marshal credentials")
	}

	return flowId[:], credentialBytes, nil
}

func (s *Service) PublicKeyAttestation(ctx context.Context, user *model.User, flowId []byte, credential *protocol.CredentialCreationResponse) ([]byte, error) {
	parsed, err := credential.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "can not parse credential")
	}

	flowIdBytes, err := uuid.FromBytes(flowId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse flowId")
	}

	user.SetWebAuthnID(flowIdBytes[:])

	data, err := s.webAuthnRepository.Get(ctx, flowIdBytes.String())
	if err != nil {
		return nil, errors.Wrap(err, "can not get session data")
	}

	createdCredential, err := s.webAuthn.CreateCredential(user, *data, parsed)
	if err != nil {
		return nil, errors.Wrap(err, "can not create credential")
	}

	var requestedFrom string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("grpcgateway-user-agent")
		if len(values) > 0 {
			userAgentInfo := useragent.Parse(values[0])
			requestedFrom = fmt.Sprintf("%s %s, %s %s", userAgentInfo.OS, userAgentInfo.OSVersion, userAgentInfo.Name, userAgentInfo.Version)
		}
	}

	if err := s.webAuthnRepository.AddUserCredential(
		ctx,
		user.Id.String(),
		flowIdBytes[:],
		createdCredential,
		requestedFrom,
	); err != nil {
		return nil, errors.Wrap(err, "can not add credential")
	}
	return nil, nil
}

func (s *Service) RequestPublicKeyAssertion(ctx context.Context) ([]byte, []byte, error) {
	login, data, err := s.webAuthn.BeginDiscoverableLogin()
	if err != nil {
		return nil, nil, errors.Wrap(err, "can not begin login")
	}
	flowId := uuid.New()
	if err := s.webAuthnRepository.Set(ctx, flowId.String(), data); err != nil {
		return nil, nil, errors.Wrap(err, "can not save session")
	}
	credentialBytes, err := json.Marshal(login.Response)
	if err != nil {
		return nil, nil, err
	}
	return flowId[:], credentialBytes, nil
}

func (s *Service) PublicKeyAssertion(ctx context.Context, flowId []byte, credential *protocol.CredentialAssertionResponse) (*model.User, error) {
	parsed, err := credential.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "can not parse credential")
	}

	flowIdBytes, err := uuid.FromBytes(flowId)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse flowId")
	}

	data, err := s.webAuthnRepository.Get(ctx, flowIdBytes.String())
	if err != nil {
		return nil, errors.Wrap(err, "can not get session data")
	}

	var user *model.User
	discoverableLogin, err := s.webAuthn.ValidateDiscoverableLogin(func(_, userHandle []byte) (webauthn.User, error) {
		userId, err := s.webAuthnRepository.GetUserIdByCredentialId(ctx, userHandle)
		if err != nil {
			return nil, err
		}

		user, err = s.userRepository.GetByID(ctx, userId)
		if err != nil {
			return nil, err
		}

		user.SetWebAuthnID(userHandle)

		userPublicKeys, err := s.webAuthnRepository.GetUserPublicKeys(ctx, user.Id.String())
		if err != nil {
			return nil, errors.Wrap(err, "found no credentials for user")
		}

		var credentials []webauthn.Credential
		for _, key := range userPublicKeys {
			credentials = append(credentials, key.Credential)
		}
		user.SetWebAuthnCredentials(credentials)
		return user, nil
	}, *data, parsed)
	if err != nil {
		return nil, errors.Wrap(err, "can not validate discoverable login")
	}
	log.Printf("discoverableLogin: %#v", discoverableLogin)
	if err := s.webAuthnRepository.UpdateUserCredential(ctx, user.Id.String(), user.WebAuthnID()[:], discoverableLogin); err != nil {
		return nil, errors.Wrap(err, "can not update credential")
	}
	return user, nil
}

func (s *Service) GetPublicKeys(ctx context.Context, user *model.User) ([]model.PublicKey, error) {
	return s.webAuthnRepository.GetUserPublicKeys(ctx, user.Id.String())
}
func (s *Service) DeletePublicKey(ctx context.Context, id []byte) error {
	return s.webAuthnRepository.DeleteCredentialById(ctx, id)
}
