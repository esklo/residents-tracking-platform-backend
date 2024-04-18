package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func (s *Service) CreateToken(ctx context.Context, id uuid.UUID) (string, error) {
	date := time.Now()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "app",
		"sub": id.String(),
		"exp": date.Add(24 * time.Hour).Unix(),
		"iat": date.Unix(),
	}).SignedString([]byte(s.appConfig.JwtSecret()))
}
