package jwt

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pallavi/sendit-api/pkg/application"
)

// Middleware returns an Echo Middleware function for JSON web token auth
func Middleware(app *application.App) echo.MiddlewareFunc {
	return middleware.JWT([]byte(app.Config.JWTSecret))
}

// SkipMiddleware returns an Echo Middleware function that can be passed in
// to routes that do not require authentication
func SkipMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			return true
		},
	})
}
