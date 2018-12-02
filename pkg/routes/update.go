package routes

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	validator "gopkg.in/go-playground/validator.v9"
)

// UpdateParams are the payload for updating a route
type UpdateParams struct {
	LocationID int    `form:"location" validate:"omitempty"`
	Name       string `form:"name" validate:"omitempty,ascii,min=1,max=100"`
	Type       string `form:"type" validate:"omitempty,oneof=boulder toprope lead"`
	Grade      string `form:"grade" validate:"omitempty,grade"`
}

// ValidateUpdateParams is a custom struct-level validator that
// requires type and grade to be passed in together
func ValidateUpdateParams(sl validator.StructLevel) {
	route := sl.Current().Interface().(UpdateParams)
	if route.Type != "" && route.Grade == "" {
		sl.ReportError(route.Grade, "Grade", "grade", "required", "")
	}
	if route.Grade != "" && route.Type == "" {
		sl.ReportError(route.Type, "Type", "type", "required", "")
	}
}

func (h *handler) update(c echo.Context) error {
	payload := UpdateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	rid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("route")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}

	route := Route{ID: rid}
	columns := []string{"date_modified"}
	if payload.LocationID != 0 {
		route.LocationID = payload.LocationID
		columns = append(columns, "location_id")
	}
	if payload.Name != "" {
		route.Name = payload.Name
		columns = append(columns, "name")
	}
	if payload.Type != "" {
		route.Type = payload.Type
		columns = append(columns, "type")
	}
	if payload.Grade != "" {
		route.Grade = payload.Grade
		columns = append(columns, "grade")
	}

	result, err := h.app.DB.Model(&route).
		Column(columns...).
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
	return h.retrieve(c)
}
