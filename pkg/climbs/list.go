package climbs

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type listParams struct {
	SessionID int   `query:"session" validate:"omitempty"`
	RouteID   int   `query:"route" validate:"omitempty"`
	Sent      *bool `query:"sent" validate:"omitempty"`
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
	climbs := []models.Climb{}
	query := h.app.DB.Model(&climbs).
		Where("deleted = false")
	if params.SessionID != 0 {
		session := models.Session{ID: params.SessionID}
		if !session.BelongsToUser(claims.ID, h.app.DB) {
			return errors.NotFound("session")
		}
		query = query.Where("session_id = ?", params.SessionID)
	}
	if params.RouteID != 0 {
		route := models.Route{ID: params.RouteID}
		if !route.BelongsToUser(claims.ID, h.app.DB) {
			return errors.NotFound("route")
		}
		query = query.Where("route_id = ?", params.RouteID)
	}
	if params.Sent != nil {
		query = query.Where("sent = ?", *params.Sent)
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("climb", err.Error())
	}
	return c.JSON(http.StatusOK, climbs)
}
