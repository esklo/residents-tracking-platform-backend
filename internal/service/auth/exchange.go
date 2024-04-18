package auth

import (
	"context"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (s *Service) ExchangeToken(ctx context.Context, tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrorUnexpectedSigningMethod
		}
		return []byte(s.appConfig.JwtSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, err := claims.GetSubject()
		if err != nil {
			return nil, err
		}
		user, err := s.userRepository.GetByID(ctx, sub)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, model.ErrorInvalidToken
}

func (s *Service) ExchangeTokenFromContext(ctx context.Context) (*model.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, model.ErrorMetadataIsEmpty
	}
	values := md.Get("authorization")
	if len(values) < 1 {
		return nil, model.ErrorAuthHeaderFormatIsInvalid
	}

	token := strings.Fields(values[0])

	switch strings.ToLower(token[0]) {
	case "bearer":
		return s.ExchangeToken(ctx, token[1])
	default:
		return nil, model.ErrorUnknownAuthType
	}
}
