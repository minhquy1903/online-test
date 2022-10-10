package http

import (
	"net/http"

	"github.com/minhquy1903/online-test/auth-service/internal/auth"
	"github.com/minhquy1903/online-test/auth-service/internal/auth/presenter"
	"github.com/minhquy1903/online-test/auth-service/utils"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	useCase auth.UseCase
}

func NewAuthHandler(useCase auth.UseCase) auth.Handler {
	return &authHandler{
		useCase: useCase,
	}
}

func (h *authHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.RegisterRequest{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		user, err := h.useCase.Register(c.Request().Context(), input.Username, input.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, presenter.RegisterResponse{Id: user.Id, Name: user.Name})
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