package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
    svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
    return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx echo.Context) {
    authorization := ctx.Request().Header.Get("authorization")

    if authorization == "" {
        ctx.Response()
        return
    }

    token := strings.Split(authorization, "Bearer ")

    if len(token) < 2 {
        ctx.Response()
        return
    }

    res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest {
        Token: token[1],
    })

    if err != nil || res.Status != http.StatusOK {
        ctx.Response()
        return
    }

    ctx.Set("userId", res.UserId)
}