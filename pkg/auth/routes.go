package auth

import (
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/application"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type handler struct {
	app *application.App
}

// RegisterRoutes registers routes onto the Echo router
func RegisterRoutes(e *echo.Echo, app *application.App) {
	h := handler{app}

	g := e.Group("/auth")
	g.POST("/login", h.login, jwt.SkipMiddleware())
	g.GET("/refresh", h.refresh, jwt.Middleware(app))
}
