package climbs

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) retrieve(c echo.Context) error {
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("climb")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	climb := models.Climb{ID: cid}
	err = h.app.DB.Model(&climb).
		WherePK().
		Select()
	if err != nil {
		return errors.DatabaseError("climb", err.Error())
	}
	if !climb.BelongsToUser(claims.ID, h.app.DB) {
		return errors.NotFound("climb")
	}
	return c.JSON(http.StatusOK, climb)
}
