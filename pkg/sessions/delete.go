package sessions

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pallavi/sendit-api/pkg/jwt"
)

func (h *handler) delete(c echo.Context) error {
	sid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		return err
	}
	session := Session{
		ID:      sid,
		Deleted: true,
	}
	result, err := h.app.DB.Model(&session).
		Column("date_modified", "deleted").
		WherePK().
		Where("user_id = ?", claims.ID).
		Where("deleted = false").
		Update()
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return c.JSON(http.StatusNotFound, "session not found")
	}
	return c.JSON(http.StatusOK, session)
}
