package users

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

	g := e.Group("/users")
	g.POST("", h.create, jwt.SkipMiddleware())
	g.GET("/:id", h.retrieve, jwt.Middleware(app))
	g.PUT("/:id", h.update, jwt.Middleware(app))
}
