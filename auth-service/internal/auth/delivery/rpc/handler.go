package rpc

import (
	"context"
	"net/http"

	"github.com/minhquy1903/online-test/auth-service/internal/auth/delivery/pb"
)

type authHandler struct {
	authService pb.AuthServiceServer
}

// func NewAuthHandler(service *service.AuthService) pb.AuthServiceServer {
// 	return &authHandler{
// 		authService: service,
// 	}
// }

func (h *authHandler) Register(ctx context.Context,req *pb.RegisterRequest) *pb.RegisterResponse {
	// fmt.Println("dsadsa", req.Email, req.Password)
	h.authService.Register(ctx, req.Email)
	return &pb.RegisterResponse{
        Status: http.StatusCreated,
    }
}

// func (h *authHandler) SignIn() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		input := &presenter.LoginInput{}
// 		if err := utils.ReadRequest(c, input); err != nil {
// 			return echo.NewHTTPError(http.StatusBadRequest)
// 		}
// 		token, err := h.useCase.SignIn(c.Request().Context(), input.Username, input.Password)
// 		if err != nil {
// 			if err == auth.ErrUserNotFound {
// 				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
// 			}
// 			if err == auth.ErrWrongPassword {
// 				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
// 			}
// 			return echo.NewHTTPError(http.StatusInternalServerError)
// 		}
// 		return c.JSON(http.StatusOK, presenter.LogInResponse{Token: token})
// 	}
// }