package routes

import (
	"reflect"
	"regexp"
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/pallavi/sendit-api/pkg/locations"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	boulderingGrade = `V(B|0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15)(\+|-)?`
	climbingGrade   = `5.(5|6|7|8|9|10|11|12|13|14|15)(a|b|c|d)?(\+|-)?`
)

// Route represents the database model for a route
type Route struct {
	tableName struct{} `sql:"routes"`

	ID           int                `sql:"id,pk" json:"id"`
	UserID       int                `sql:"user_id,notnull" json:"user"`
	LocationID   int                `sql:"location_id" json:"-"`
	Location     locations.Location `json:"location"`
	Name         string             `sql:"name,notnull" json:"name"`
	Type         string             `sql:"type,notnull" json:"type"`
	Grade        string             `sql:"grade,notnull" json:"grade"`
	Tags         []string           `sql:"tags,array" json:"tags"`
	Deleted      bool               `sql:"deleted" json:"deleted,omitempty"`
	DateCreated  time.Time          `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time          `sql:"date_modified,notnull" json:"date_modified"`
}

// ValidateGrade does cross-field validation of grade and type
func ValidateGrade(fl validator.FieldLevel) bool {
	rType := reflect.Indirect(fl.Parent()).FieldByName("Type").String()
	rGrade := fl.Field().String()
	var regex string
	if rType == "boulder" {
		regex = boulderingGrade
	} else {
		regex = climbingGrade
	}
	match, _ := regexp.MatchString(regex, rGrade)
	return match
}

// BeforeInsert sets date fields
func (r *Route) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if r.DateCreated.IsZero() {
		r.DateCreated = now
	}
	r.DateModified = now
	return nil
}

// BeforeUpdate sets the date_modified field
func (r *Route) BeforeUpdate(db orm.DB) error {
	r.DateModified = time.Now()
	return nil
}

// BelongsToUser checks if the route belongs to the given user
func (r *Route) BelongsToUser(uid int, db orm.DB) bool {
	err := db.Model(r).
		WherePK().
		Where("user_id = ?", uid).
		Select()
	if err != nil {
		return false
	}
	return true
}
