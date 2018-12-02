package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type listParams struct {
	LocationID int    `query:"location" validate:"omitempty"`
	Type       string `query:"type" validate:"omitempty,oneof=boulder toprope lead"`
	Grade      string `query:"grade" validate:"omitempty,grade"`
}

func (h *handler) list(c echo.Context) error {
	params := listParams{}
	err := c.Bind(&params)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	routes := []Route{}
	query := h.app.DB.Model(&routes).
		Where("user_id = ?", claims.ID).
		Where("deleted = false")
	if params.LocationID != 0 {
		query = query.Where("location_id = ?", params.LocationID)
	}
	if params.Type != "" {
		query = query.Where("type = ?", params.Type)
	}
	if params.Grade != "" {
		query = query.Where("grade = ?", params.Grade)
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	return c.JSON(http.StatusOK, routes)
}
