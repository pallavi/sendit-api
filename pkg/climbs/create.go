package climbs

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/routes"
	"github.com/pallavi/sendit-api/pkg/sessions"
)

type createParams struct {
	SessionID int    `form:"session" validate:"required"`
	RouteID   int    `form:"route" validate:"required"`
	Attempts  int    `form:"attempts" validate:"required,min=1"`
	Sent      *bool  `form:"sent" validate:"required"`
	Notes     string `form:"notes" validate:"omitempty,max=5000"`
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
	session := sessions.Session{ID: payload.SessionID}
	if !session.BelongsToUser(claims.ID, h.app.DB) {
		return errors.NotFound("session")
	}
	route := routes.Route{ID: payload.RouteID}
	if !route.BelongsToUser(claims.ID, h.app.DB) {
		return errors.NotFound("route")
	}
	climb := Climb{
		SessionID: payload.SessionID,
		RouteID:   payload.RouteID,
		Attempts:  payload.Attempts,
		Sent:      *payload.Sent,
		Notes:     payload.Notes,
	}
	err = h.app.DB.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model(&climb).Insert()
		return err
	})
	if err != nil {
		return errors.DatabaseError("climb", err.Error())
	}
	return c.JSON(http.StatusOK, climb)
}
