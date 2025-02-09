package auth

import (
	"context"
	"encoding/json"
	proto "github.com/esklo/residents-tracking-platform-backend/gen/proto/auth"
	"github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/pkg/errors"
)

func (i Implementation) RequestPublicKeyAttestation(ctx context.Context, _ *empty.Empty) (*proto.PublicKeyCredentialsResponse, error) {
	i.logger.Debug("auth.RequestPublicKeyAttestation request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	flowId, credentials, err := i.authService.RequestPublicKeyAttestation(ctx, user)
	if err != nil {
		return nil, err
	}
	return &proto.PublicKeyCredentialsResponse{
		Credentials: credentials,
		FlowId:      flowId,
	}, nil
}

func (i Implementation) PublicKeyAttestation(ctx context.Context, req *proto.PublicKeyCredentialsRequest) (*empty.Empty, error) {
	i.logger.Debug("auth.PublicKeyAttestation request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	var ccr protocol.CredentialCreationResponse
	if err := json.Unmarshal([]byte(req.Credentials), &ccr); err != nil {
		return nil, errors.Wrap(err, "can not unmarshal credentials")
	}

	if _, err := i.authService.PublicKeyAttestation(ctx, user, req.FlowId, &ccr); err != nil {
		return nil, errors.Wrap(err, "can not submit webauthn credentials")
	}
	return &empty.Empty{}, nil
}

func (i Implementation) RequestPublicKeyAssertion(ctx context.Context, _ *empty.Empty) (*proto.PublicKeyCredentialsResponse, error) {
	i.logger.Debug("auth.RequestPublicKeyAssertion request")
	flowId, credentials, err := i.authService.RequestPublicKeyAssertion(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.PublicKeyCredentialsResponse{
		Credentials: credentials,
		FlowId:      flowId,
	}, nil
}

func (i Implementation) PublicKeyAssertion(ctx context.Context, req *proto.PublicKeyCredentialsRequest) (*proto.LoginResponse, error) {
	i.logger.Debug("auth.PublicKeyAssertion request")
	var car protocol.CredentialAssertionResponse
	if err := json.Unmarshal([]byte(req.Credentials), &car); err != nil {
		return nil, errors.Wrap(err, "can not unmarshal credentials")
	}
	user, err := i.authService.PublicKeyAssertion(ctx, req.FlowId, &car)
	if err != nil {
		return nil, err
	}

	token, err := i.authService.CreateToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	protoUser, err := user.ToProto()
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{
		Token: token,
		User:  protoUser,
	}, nil
}

func (i Implementation) GetPublicKeys(ctx context.Context, req *empty.Empty) (*proto.GetPublicKeysResponse, error) {
	i.logger.Debug("auth.GetPublicKeys request")
	user, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}
	var keys []*proto.Key
	publicKeys, err := i.authService.GetPublicKeys(ctx, user)
	if err != nil {
		return nil, err
	}
	for _, key := range publicKeys {
		protoKey, err := key.ToProto()
		if err != nil {
			return nil, err
		}
		keys = append(keys, protoKey)
	}
	return &proto.GetPublicKeysResponse{
		Keys: keys,
	}, nil
}

func (i Implementation) DeletePublicKey(ctx context.Context, key *proto.Key) (*empty.Empty, error) {
	i.logger.Debug("auth.DeletePublicKey request")
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	if err := i.authService.DeletePublicKey(ctx, key.Id); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
