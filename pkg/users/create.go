package users

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
)

// CreateParams contains a username and password
type CreateParams struct {
	Username string `form:"username" validate:"required,alphanum,min=1,max=100"`
	Password string `form:"password" validate:"required,min=8"`
}

func (h *handler) create(c echo.Context) error {
	payload := CreateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	user := User{
		Username: payload.Username,
		Password: payload.Password,
	}
	err = h.app.DB.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model(&user).Insert()
		return err
	})
	if err != nil {
		return errors.DatabaseError("user", err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
