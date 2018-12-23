package locations

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Location represents the database model for a user
type Location struct {
	tableName struct{} `sql:"locations"`

	ID           int       `sql:"id,pk" json:"id"`
	UserID       int       `sql:"user_id,notnull" json:"user"`
	Name         string    `sql:"name,notnull" json:"name"`
	Outdoor      bool      `sql:"outdoor,notnull" json:"outdoor"`
	Deleted      bool      `sql:"deleted" json:"deleted,omitempty"`
	DateCreated  time.Time `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time `sql:"date_modified,notnull" json:"date_modified"`
}

// BeforeInsert sets date fields
func (l *Location) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if l.DateCreated.IsZero() {
		l.DateCreated = now
	}
	l.DateModified = now
	return nil
}

// BeforeUpdate sets the date_modified field
func (l *Location) BeforeUpdate(db orm.DB) error {
	l.DateModified = time.Now()
	return nil
}

// BelongsToUser checks if the location belongs to the given user
func (l *Location) BelongsToUser(uid int, db orm.DB) bool {
	err := db.Model(l).
		WherePK().
		Where("user_id = ?", uid).
		Select()
	if err != nil {
		return false
	}
	return true
}
