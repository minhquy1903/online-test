package auth

import (
	"context"

	"github.com/minhquy1903/online-test/internal/auth/domain"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) (*domain.User, error)
	// SignIn(ctx context.Context, username, password string) (string, error)
	// ParseToken(ctx context.Context, accessToken string) (string, error)
}