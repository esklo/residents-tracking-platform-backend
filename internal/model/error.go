package model

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrorInvalidCredentials        = errors.New("invalid credentials")
	ErrorNotFound                  = errors.New("not found")
	ErrorMetadataIsEmpty           = errors.New("metadata is empty")
	ErrorAuthHeaderFormatIsInvalid = errors.New("auth header format is invalid")
	ErrorUnknownAuthType           = errors.New("unknown auth type")
	ErrorCanNotParseQuery          = errors.New("can not parse query")
	ErrorUnauthenticated           = status.Error(codes.Unauthenticated, "unauthenticated")
	ErrorModelIsEmpty              = status.Error(codes.Internal, "model is empty")
	ErrorInvalidToken              = status.Error(codes.Unauthenticated, "invalid token")
	ErrorUnexpectedSigningMethod   = errors.New("unexpected signing method")
	ErrorPermissionDenied          = status.Error(codes.PermissionDenied, "permission denied")
	ErrorCurrentPasswordIsInvalid  = status.Error(codes.PermissionDenied, "current password is invalid")
	ErrorInvalidUUID               = status.Error(codes.InvalidArgument, "invalid uuid")
)
