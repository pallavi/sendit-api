package locations

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type listParams struct {
	Outdoor *bool `query:"outdoor" validate:"omitempty"`
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
	locations := []models.Location{}
	query := h.app.DB.Model(&locations).
		Where("user_id = ?", claims.ID).
		Where("deleted = false")
	if params.Outdoor != nil {
		query = query.Where("outdoor = ?", *params.Outdoor)
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("location", err.Error())
	}
	return c.JSON(http.StatusOK, locations)
}
