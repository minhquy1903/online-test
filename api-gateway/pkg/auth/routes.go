package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/online-test/api-gateway/pkg/config"
)


func RegisterRoutes(r *echo.Group, c *config.Config) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }
    
    routes := r.Group("/auth")
    routes.POST("/register", svc.Register)
    routes.POST("/login", svc.Login)

    return svc
}

