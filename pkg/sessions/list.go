package sessions

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type listParams struct {
	LocationID int `query:"location" validate:"omitempty"`
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
	sessions := []models.Session{}
	query := h.app.DB.Model(&sessions).
		Where("user_id = ?", claims.ID).
		Where("deleted = false")
	if params.LocationID != 0 {
		query = query.Where("location_id = ?", params.LocationID)
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("session", err.Error())
	}
	return c.JSON(http.StatusOK, sessions)
}
