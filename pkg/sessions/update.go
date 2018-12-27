package sessions

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/binder"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type updateParams struct {
	StartTime *binder.Timestamp `form:"start" validate:"omitempty"`
	EndTime   *binder.Timestamp `form:"end" validate:"omitempty,gtefield=StartTime"`
}

func (h *handler) update(c echo.Context) error {
	payload := updateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	sid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.InvalidID("session")
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}

	session := models.Session{ID: sid}
	columns := []string{"date_modified"}
	if payload.StartTime != nil {
		session.StartTime = time.Time(*payload.StartTime)
		columns = append(columns, "start_time")
	}
	if payload.EndTime != nil {
		end := time.Time(*payload.EndTime)
		session.EndTime = &end
		columns = append(columns, "end_time")
	}

	result, err := h.app.DB.Model(&session).
		Column(columns...).
		WherePK().
		Where("user_id = ?", claims.ID).
		Where("deleted = false").
		Update()
	if err != nil {
		return errors.DatabaseError("session", err.Error())
	}
	if result.RowsAffected() == 0 {
		return errors.NotFound("session")
	}
	return h.retrieve(c)
}
