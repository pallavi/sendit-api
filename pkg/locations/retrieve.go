package locations

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) retrieve(c echo.Context) error {
	lid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("location")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	location := models.Location{ID: lid}
	err = h.app.DB.Model(&location).
		WherePK().
		Where("user_id = ?", claims.ID).
		Select()
	if err != nil {
		return errors.DatabaseError("location", err.Error())
	}
	return c.JSON(http.StatusOK, location)
}
