package users

import (
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

type updateParams struct {
	Username *string `form:"username" validate:"omitempty,alphanum,min=1,max=100"`
	Password *string `form:"password" validate:"omitempty,min=8"`
}

func (h *handler) update(c echo.Context) error {
	payload := updateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return errors.BadJWTClaims(err.Error())
	}

	user := models.User{ID: claims.ID}
	columns := []string{"date_modified"}
	if payload.Username != nil {
		user.Username = *payload.Username
		columns = append(columns, "username")
	}
	if payload.Password != nil {
		user.Password = *payload.Password
		columns = append(columns, "password")
	}

	_, err = h.app.DB.Model(&user).
		Column(columns...).
		WherePK().
		Update()
	if err != nil {
		return errors.DatabaseError("user", err.Error())
	}
	return h.retrieve(c)
}
