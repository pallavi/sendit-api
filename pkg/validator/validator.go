package validator

import (
	"github.com/pallavi/sendit-api/pkg/routes"
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator conforms with Echo's Validator type
// https://echo.labstack.com/guide/request#validate-data
type CustomValidator struct {
	validator *validator.Validate
}

// New constructor for creating a new custom validator
func New() *CustomValidator {
	validator := validator.New()
	validator.RegisterValidation("grade", routes.ValidateGrade)
	validator.RegisterStructValidation(routes.ValidateUpdateParams, routes.UpdateParams{})
	return &CustomValidator{validator}
}

// Validate validates a struct using the third party validator's Validate function
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
