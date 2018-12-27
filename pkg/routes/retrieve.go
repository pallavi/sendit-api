package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) retrieve(c echo.Context) error {
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("route")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	route := models.Route{ID: rid}
	err = h.app.DB.Model(&route).
		WherePK().
		Where("route.user_id = ?", claims.ID).
		Relation("Location").
		Select()
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	return c.JSON(http.StatusOK, route)
}
