package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/minhquy1903/online-test/auth-service/config"
	"github.com/minhquy1903/online-test/auth-service/model"
	"github.com/minhquy1903/online-test/auth-service/pb"
	"github.com/minhquy1903/online-test/auth-service/utils"
	"gorm.io/gorm"
)

type Handler struct {
    db   *gorm.DB
    cfg	 *config.Config
}

func NewHandler(db *gorm.DB, cfg *config.Config) *Handler {
	return &Handler{
		db: db,
		cfg: cfg,
	}
}

func (a *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// var user model.User

    fmt.Println("yeahhhhh", req.Password)

    // if result := a.db.Where(&model.User{Email: req.Email}).First(&user); result.Error == nil {
    //     return &pb.RegisterResponse{
    //         Status: http.StatusConflict,
    //         Error:  "E-Mail already exists",
    //     }, nil
    // }

    // user.Email = req.Email
    // user.Password = utils.HashPassword(req.Password)

    // a.db.Create(&user)

    return &pb.RegisterResponse{
        Status: http.StatusCreated,
    }, nil
}

func (a *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User

    if result := a.db.Where(&model.User{Email: req.Email}).First(&user); result.Error != nil {
        return &pb.LoginResponse{
            Status: http.StatusNotFound,
            Error:  "User not found",
        }, nil
    }

    match := utils.CheckPasswordHash(req.Password, user.Password)

    if !match {
        return &pb.LoginResponse{
            Status: http.StatusNotFound,
            Error:  "User not found",
        }, nil
    }

    token := ".GenerateToken(user)"

    return &pb.LoginResponse{
        Status: http.StatusOK,
        Token:  token,
    }, nil
}

// func (a *AuthService) ParseToken(ctx context.Context, accessToken string) (string, error) {
// 	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return a.cfg.Server.ReadTimeout, nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
// 		return claims.UserId, nil
// 	}

// 	return "", nil
// }
