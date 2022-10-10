package auth

import (
	"context"

	"github.com/minhquy1903/online-test/auth-service/internal/auth/domain"
)

const CtxUserKey = "userId"

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserById(ctx context.Context, userId string) (*domain.User, error)
}