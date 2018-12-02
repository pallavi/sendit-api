package sessions

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

	g := e.Group("/sessions", jwt.Middleware(app))
	g.POST("", h.create)
	g.GET("", h.list)
	g.GET("/:id", h.retrieve)
	g.PUT("/:id", h.update)
	g.DELETE("/:id", h.delete)
}
