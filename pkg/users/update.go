package users

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
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
	uid := c.Param("id")
	claims, err := jwt.GetClaims(c)
	if uid != strconv.Itoa(claims.ID) {
		return errors.NotFound("user")
	}

	user := User{ID: claims.ID}
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
