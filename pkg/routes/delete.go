package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

func (h *handler) delete(c echo.Context) error {
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("route")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	route := Route{
		ID:      rid,
		Deleted: true,
	}
	result, err := h.app.DB.Model(&route).
		Column("date_modified", "deleted").
		WherePK().
		Where("user_id = ?", claims.ID).
		Where("deleted = false").
		Update()
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	if result.RowsAffected() == 0 {
		return errors.NotFound("route")
	}
	return c.JSON(http.StatusOK, route)
}
