package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

const (
	pgDuplicateKeyMsg = "ERROR #23505"
	pgNoRowsFoundMsg  = "no rows in result set"
)

// DatabaseError returns a 500 error
func DatabaseError(resource, msg string) error {
	if strings.Contains(msg, pgDuplicateKeyMsg) {
		return Duplicate(resource)
	}
	if strings.Contains(msg, pgNoRowsFoundMsg) {
		return NotFound(resource)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("database error: %s", msg))
}

// NotFound returns a 404 error with a message indicating the given resource
func NotFound(resource string) error {
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}

// BadRequest returns a 400 error when request fails validation
func BadRequest(msg string) error {
	errMsg := strings.Split(msg, "Error:")
	if len(errMsg) > 1 {
		msg = errMsg[1]
	}
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("validation failed: %s", msg))
}

// InvalidID returns a 400 error on invalid IDs in the path
func InvalidID(resource string) error {
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid %s id", resource))
}

// Duplicate returns a 400 error when unique constraint is violated
func Duplicate(resource string) error {
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s already exists", resource))
}

// BadCredentials returns a 401 error on failed login
func BadCredentials() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "username or password is incorrect")
}

// BadJWTClaims returns a 401 error when JWT claims cannot be found on the Echo context
func BadJWTClaims(msg string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("error getting jwt claims: %s", msg))
}
