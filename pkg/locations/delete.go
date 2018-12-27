package locations

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) delete(c echo.Context) error {
	lid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("location")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	location := models.Location{
		ID:      lid,
		Deleted: true,
	}
	result, err := h.app.DB.Model(&location).
		Column("date_modified", "deleted").
		WherePK().
		Where("user_id = ?", claims.ID).
		Where("deleted = false").
		Update()
	if err != nil {
		return errors.DatabaseError("location", err.Error())
	}
	if result.RowsAffected() == 0 {
		return errors.NotFound("location")
	}
	return c.JSON(http.StatusOK, location)
}
