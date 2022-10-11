package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/minhquy1903/online-test/auth-service/internal/auth"
	"github.com/minhquy1903/online-test/auth-service/internal/auth/model"
	"gorm.io/gorm"
)

type AuthClaims struct {
	jwt.StandardClaims
	Email string `json:"username"`
	UserId   string `json:"userId"`
}

type AuthService struct {
    db   *gorm.DB
    cfg	 *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		db: db,
		cfg: cfg,
	}
}

func (a *AuthService) Register(ctx context.Context, email string) (*model.User, error) {
	return nil,nil
}

func (a *AuthService) Login(ctx context.Context, email string) (string, error) {
	user := model.User{}

	claims := AuthClaims{
		Email: user.Email,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "go-todos",
			ExpiresAt: time.Now().Add(a.cfg.Server.ReadTimeout).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.cfg.Server.ReadTimeout)
}

func (a *AuthService) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.cfg.Server.ReadTimeout, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", auth.ErrInvalidAccessToken
}
