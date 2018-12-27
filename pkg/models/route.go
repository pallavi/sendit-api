package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Route represents the database model for a route
type Route struct {
	tableName struct{} `sql:"routes"`

	ID           int       `sql:"id,pk" json:"id"`
	UserID       int       `sql:"user_id,notnull" json:"user"`
	LocationID   int       `sql:"location_id" json:"-"`
	Location     Location  `json:"location"`
	Name         string    `sql:"name,notnull" json:"name"`
	Type         string    `sql:"type,notnull" json:"type"`
	Grade        string    `sql:"grade,notnull" json:"grade"`
	Tags         []string  `sql:"tags,array" json:"tags"`
	Deleted      bool      `sql:"deleted" json:"deleted,omitempty"`
	DateCreated  time.Time `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time `sql:"date_modified,notnull" json:"date_modified"`
	Climbs       []Climb   `json:"climbs,omitempty"`
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
