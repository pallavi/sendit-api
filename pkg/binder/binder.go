package binder

import (
	"time"

	"github.com/labstack/echo"
)

// CustomBinder struct
type CustomBinder struct {
	binder *echo.DefaultBinder
}

// New constructor for creating a new custom binder
func New() *CustomBinder {
	return &CustomBinder{}
}

// Bind binds and validates payloads using the default Echo binder
// and the validator set on Echo's Context
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	err := cb.binder.Bind(i, c)
	if err != nil {
		return err
	}
	return c.Validate(i)
}

// Timestamp implements the Echo#BindUnmarshaler interface for a custom type
type Timestamp time.Time

// UnmarshalParam converts a string into a Timestamp
func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}
