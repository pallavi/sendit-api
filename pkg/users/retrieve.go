package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

func (h *handler) retrieve(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("user")
	}
	claims, err := jwt.GetClaims(c)
	if uid != claims.ID {
		return errors.NotFound("user")
	}
	user := &User{ID: uid}
	err = h.app.DB.Select(user)
	if err != nil {
		return errors.DatabaseError("user", err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
