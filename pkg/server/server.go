package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pallavi/sendit-api/pkg/application"
	"github.com/pallavi/sendit-api/pkg/auth"
	"github.com/pallavi/sendit-api/pkg/binder"
	"github.com/pallavi/sendit-api/pkg/climbs"
	"github.com/pallavi/sendit-api/pkg/locations"
	"github.com/pallavi/sendit-api/pkg/routes"
	"github.com/pallavi/sendit-api/pkg/sessions"
	"github.com/pallavi/sendit-api/pkg/users"
	"github.com/pallavi/sendit-api/pkg/validator"
)

// New returns an Echo server
func New(app *application.App) *http.Server {
	e := echo.New()

	e.Validator = validator.New()
	e.Binder = binder.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	auth.RegisterRoutes(e, app)
	users.RegisterRoutes(e, app)
	locations.RegisterRoutes(e, app)
	routes.RegisterRoutes(e, app)
	sessions.RegisterRoutes(e, app)
	climbs.RegisterRoutes(e, app)

	return e.Server
}
