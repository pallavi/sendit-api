package climbs

import (
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/pallavi/sendit-api/pkg/sessions"
)

// Climb represents the database model for a climb
type Climb struct {
	tableName struct{} `sql:"climbs"`

	ID           int       `sql:"id,pk" json:"id"`
	SessionID    int       `sql:"session_id,notnull" json:"session"`
	RouteID      int       `sql:"route_id,notnull" json:"route"`
	Attempts     int       `sql:"attempts,notnull" json:"attempts"`
	Sent         bool      `sql:"sent,notnull" json:"sent"`
	Deleted      bool      `sql:"deleted" json:"deleted,omitempty"`
	DateCreated  time.Time `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time `sql:"date_modified,notnull" json:"date_modified"`
}

// BeforeInsert sets date fields
func (c *Climb) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if c.DateCreated.IsZero() {
		c.DateCreated = now
	}
	c.DateModified = now
	return nil
}

// BeforeUpdate sets the date_modified field
func (c *Climb) BeforeUpdate(db orm.DB) error {
	c.DateModified = time.Now()
	return nil
}

// BelongsToUser checks if the climb belongs to the given user
func (c *Climb) BelongsToUser(uid int, db orm.DB) bool {
	err := db.Model(&c).
		WherePK().
		Select()
	if err != nil {
		return false
	}
	session := sessions.Session{ID: c.SessionID}
	return session.BelongsToUser(uid, db)
}
