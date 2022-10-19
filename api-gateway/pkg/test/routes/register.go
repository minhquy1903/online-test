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
    Name     string `json:"name"`
    Type     int32 `json:"type"`
}

func Register(ctx echo.Context, c pb.AuthServiceClient) {
    b := RegisterRequestBody{}

    if err := ctx.Bind(&b); err != nil {
        ctx.String(http.StatusBadRequest, "Bad Request1")
        return
    }

    res, err := c.Register(context.Background(), &pb.RegisterRequest{
        Email:    b.Email,
        Password: b.Password,
        Name:     b.Name,
        Type:     b.Type,
    })

    if err != nil {
        ctx.String(http.StatusBadRequest, err.Error())
        return
    }

    ctx.JSON(int(res.Status), &res)
}