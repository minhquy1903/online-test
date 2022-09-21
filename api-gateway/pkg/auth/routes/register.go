package routes

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Register(ctx echo.Context, c pb.AuthServiceClient) {
    body := RegisterRequestBody{}

    if err := ctx.Bind(&body); err != nil {
        ctx.Logger().Error(err)
        return
    }
    
    res, err := c.Register(context.Background(), &pb.RegisterRequest{
        Email:    body.Email,
        Password: body.Password,
    })

    if err != nil {
        ctx.Logger().Error(err)
        return
    }

    ctx.JSON(int(res.Status), &res)
}