package sessions

import (
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/binder"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type createParams struct {
	LocationID int               `form:"location" validate:"required"`
	StartTime  *binder.Timestamp `form:"start" validate:"omitempty"`
	EndTime    *binder.Timestamp `form:"end" validate:"omitempty,gtefield=StartTime"`
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
	session := models.Session{
		UserID:     claims.ID,
		LocationID: payload.LocationID,
	}
	if payload.StartTime != nil {
		session.StartTime = time.Time(*payload.StartTime)
	} else {
		session.StartTime = time.Now()
	}
	if payload.EndTime != nil {
		end := time.Time(*payload.EndTime)
		session.EndTime = &end
	}
	err = h.app.DB.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model(&session).Insert()
		return err
	})
	if err != nil {
		return errors.DatabaseError("session", err.Error())
	}
	err = h.app.DB.Model(&session).
		WherePK().
		Relation("Location").
		Select()
	if err != nil {
		return errors.DatabaseError("session", err.Error())
	}
	return c.JSON(http.StatusOK, session)
}
