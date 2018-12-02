package locations

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

type createParams struct {
	Name    string `form:"name" validate:"required,ascii,min=1,max=100"`
	Outdoor *bool  `form:"outdoor" validate:"required"`
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
	location := Location{
		UserID:  claims.ID,
		Name:    payload.Name,
		Outdoor: *payload.Outdoor,
	}
	err = h.app.DB.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model(&location).Insert()
		return err
	})
	if err != nil {
		return errors.DatabaseError("location", err.Error())
	}
	return c.JSON(http.StatusOK, location)
}
