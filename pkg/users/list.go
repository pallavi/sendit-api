package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
)

type listParams struct {
	Username string `query:"username" validate:"required"`
}

func (h *handler) list(c echo.Context) error {
	params := listParams{}
	err := c.Bind(&params)
	if err != nil {
		return errors.BadRequest(err.Error())
	}
	users := []User{}
	query := h.app.DB.Model(&users)
	if params.Username != "" {
		query = query.Where("username = ?", params.Username)
	}
	err = query.Order("date_created desc").Select()
	if err != nil {
		return errors.DatabaseError("user", err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
