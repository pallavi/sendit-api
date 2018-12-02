package routes

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type createParams struct {
	LocationID int    `form:"location" validate:"required"`
	Name       string `form:"name" validate:"required,ascii,min=1,max=100"`
	Type       string `form:"type" validate:"required,oneof=boulder toprope lead"`
	Grade      string `form:"grade" validate:"required,grade"`
}

func (h *handler) create(c echo.Context) error {
	payload := createParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	route := Route{
		UserID:     claims.ID,
		LocationID: payload.LocationID,
		Name:       payload.Name,
		Type:       payload.Type,
		Grade:      payload.Grade,
	}
	err = h.app.DB.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model(&route).Insert()
		return err
	})
	if err != nil {
		return errors.DatabaseError("route", err.Error())
	}
	return c.JSON(http.StatusOK, route)
}
