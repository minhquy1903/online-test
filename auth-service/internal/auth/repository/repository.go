package repository

import (
	"context"
	"strings"

	"github.com/minhquy1903/online-test/auth-service/internal/auth"
	"github.com/minhquy1903/online-test/auth-service/internal/auth/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) auth.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	result := ur.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where(&domain.User{
		Email: strings.ToLower(email),
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) GetUserById(ctx context.Context, userId string) (*domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where(&domain.User{
		Id: userId,
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}