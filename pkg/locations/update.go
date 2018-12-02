package locations

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type updateParams struct {
	Name    string `form:"name" validate:"omitempty,ascii,min=1,max=100"`
	Outdoor *bool  `form:"outdoor" validate:""`
}

func (h *handler) update(c echo.Context) error {
	payload := updateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	lid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("location")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}

	location := Location{ID: lid}
	columns := []string{"date_modified"}
	if payload.Name != "" {
		location.Name = payload.Name
		columns = append(columns, "name")
	}
	if payload.Outdoor != nil {
		location.Outdoor = *payload.Outdoor
		columns = append(columns, "outdoor")
	}

	result, err := h.app.DB.Model(&location).
		Column(columns...).
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
	return h.retrieve(c)
}
