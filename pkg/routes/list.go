package routes

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type listParams struct {
	LocationID int      `query:"location" validate:"omitempty"`
	Type       string   `query:"type" validate:"omitempty,oneof=boulder toprope lead"`
	Grade      string   `query:"grade" validate:"omitempty,grade"`
	Tags       []string `query:"tags" validate:"omitempty,dive,required,ascii,min=1,max=100"`
	Include    []string `query:"include" validate:"omitempty,dive,required,oneof=climbs"`
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
	routes := []models.Route{}
	query := h.app.DB.Model(&routes).
		Where("route.user_id = ?", claims.ID).
		Where("route.deleted = false").
		Relation("Location")
	if params.LocationID != 0 {
		query = query.Where("location_id = ?", params.LocationID)
	}
	if params.Type != "" {
		query = query.Where("type = ?", params.Type)
	}
	if params.Grade != "" {
		query = query.Where("grade = ?", params.Grade)
	}
	if params.Tags != nil {
		query = query.Where("tags @> ?", pg.Array(params.Tags))
	}
	if params.Include != nil {
		query = query.Relation("Climbs")
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	return c.JSON(http.StatusOK, routes)
}
