package validator

import (
	"reflect"
	"regexp"

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
	validator.RegisterValidation("grade", ValidateGrade)
	validator.RegisterStructValidation(routes.ValidateUpdateParams, routes.UpdateParams{})
	return &CustomValidator{validator}
}

// Validate validates a struct using the third party validator's Validate function
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// ValidateGrade does cross-field validation of grade and type
func ValidateGrade(fl validator.FieldLevel) bool {
	rType := reflect.Indirect(fl.Parent()).FieldByName("Type").String()
	rGrade := fl.Field().String()
	var regex string
	if rType == "boulder" {
		regex = `V(B|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15)(\+|-)?`
	} else {
		regex = `5.(5|6|7|8|9|10|11|12|13|14|15)(a|b|c|d)?(\+|-)?`
	}
	match, _ := regexp.MatchString(regex, rGrade)
	return match
}
