package auth

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/errors"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/users"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) login(c echo.Context) error {
	payload := users.CreateParams{}
	err := c.Bind(&payload)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	user := users.User{}
	err = h.app.DB.Model(&user).
		Where("username = ?", payload.Username).
		Select()
	if err == pg.ErrNoRows {
		return errors.BadCredentials()
	} else if err != nil {
		return errors.DatabaseError("user", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return errors.BadCredentials()
	}

	claims := jwt.Claims{
		ID:       user.ID,
		Username: user.Username,
	}
	token, err := jwt.Encode(&claims, h.app.Config.JWTSecret)
	if err != nil {
		return err
	}
	user.JWTToken = token
	return c.JSON(http.StatusOK, user)
}
