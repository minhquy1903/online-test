package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Register(ctx echo.Context, c pb.AuthServiceClient) {
    body := RegisterRequestBody{}

    if err := ctx.Bind(&body); err != nil {
        ctx.String(http.StatusBadRequest, "Bad Request1")
        return
    }

    res, err := c.Register(context.Background(), &pb.RegisterRequest{
        Email:    body.Email,
        Password: body.Password,
    })

    if err != nil {
        ctx.String(http.StatusBadRequest, "Bad Request2")
        return
    }

    ctx.JSON(int(res.Status), &res)
}