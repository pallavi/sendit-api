package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) retrieve(c echo.Context) error {
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	user := &models.User{ID: claims.ID}
	err = h.app.DB.Select(user)
	if err != nil {
		return errors.DatabaseError("user", err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
