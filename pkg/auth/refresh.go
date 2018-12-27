package auth

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/jwt"
	"github.com/pallavi/sendit-api/pkg/models"
)

func (h *handler) refresh(c echo.Context) error {
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return err
	}
	token, err := jwt.Encode(&claims, h.app.Config.JWTSecret)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.User{
		ID:       claims.ID,
		Username: claims.Username,
		JWTToken: token,
	})
}
