package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/api-gateway/pkg/auth/routes"
	"github.com/minhquy1903/api-gateway/pkg/config"
)

func RegisterRoutes(r echo.Context, c *config.Config) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }

    routes := r.Echo().Group("auth")
    routes.POST("/register", func(c echo.Context) error {return nil})
    routes.POST("/login", func(c echo.Context) error {return nil})

    return svc
}

func (svc *ServiceClient) Register(ctx echo.Context) {
    routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx echo.Context) {
    routes.Login(ctx, svc.Client)
}