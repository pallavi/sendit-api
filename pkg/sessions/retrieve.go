package sessions

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) retrieve(c echo.Context) error {
	sid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("session")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	session := models.Session{ID: sid}
	err = h.app.DB.Model(&session).
		WherePK().
		Where("user_id = ?", claims.ID).
		Select()
	if err != nil {
		return errors.DatabaseError("session", err.Error())
	}
	return c.JSON(http.StatusOK, session)
}
