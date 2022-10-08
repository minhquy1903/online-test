package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Map auth routes
func MapAuthRoutes(authGroup *echo.Group) {
	authGroup.POST("/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// authGroup.POST("/login", h.SignIn())
	// authGroup.POST("/logout", h.Logout())
}