package climbs

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type updateParams struct {
	Attempts int   `form:"attempts" validate:"omitempty,min=1"`
	Sent     *bool `form:"sent" validate:"omitempty"`
}

func (h *handler) update(c echo.Context) error {
	payload := updateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("climb")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}
	climb := Climb{ID: cid}
	if !climb.BelongsToUser(claims.ID, h.app.DB) {
		return errors.NotFound("climb")
	}

	climb = Climb{ID: cid}
	columns := []string{"date_modified"}
	if payload.Attempts != 0 {
		climb.Attempts = payload.Attempts
		columns = append(columns, "attempts")
	}
	if payload.Sent != nil {
		climb.Sent = *payload.Sent
		columns = append(columns, "sent")
	}

	result, err := h.app.DB.Model(&climb).
		Column(columns...).
		WherePK().
		Where("deleted = false").
		Update()
	if err != nil {
		return errors.DatabaseError("climb", err.Error())
	}
	if result.RowsAffected() == 0 {
		return errors.NotFound("climb")
	}
	return h.retrieve(c)
}
