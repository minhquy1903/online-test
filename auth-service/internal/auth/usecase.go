package auth

import (
	"context"

	"github.com/minhquy1903/online-test/auth-service/internal/auth/domain"
)

type UseCase interface {
	Register(ctx context.Context, username, password string) (*domain.User, error)
	// Login(ctx context.Context, username, password string) (string, error)
	// ParseToken(ctx context.Context, accessToken string) (string, error)
}