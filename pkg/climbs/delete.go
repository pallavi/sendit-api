package climbs

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) delete(c echo.Context) error {
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("climb")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	climb := models.Climb{ID: cid}
	if !climb.BelongsToUser(claims.ID, h.app.DB) {
		return errors.NotFound("climb")
	}
	climb = models.Climb{
		ID:      cid,
		Deleted: true,
	}
	result, err := h.app.DB.Model(&climb).
		Column("date_modified", "deleted").
		WherePK().
		Where("deleted = false").
		Update()
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	if result.RowsAffected() == 0 {
		return errors.NotFound("climb")
	}
	return c.JSON(http.StatusOK, climb)
}
