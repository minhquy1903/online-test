package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Login(ctx echo.Context, c pb.AuthServiceClient) {
    b := LoginRequestBody{}

    if err := ctx.Bind(&b); err != nil {
        ctx.String(http.StatusBadRequest, "Bad Request")
    }

    res, err := c.Login(context.Background(), &pb.LoginRequest{
        Email:    b.Email,
        Password: b.Password,
    })

    if err != nil {
        ctx.String(http.StatusBadGateway, "Bad Request")
    }

    ctx.JSON(http.StatusCreated, &res)
}