package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Session represents the database model for a climbing session
type Session struct {
	tableName struct{} `sql:"sessions"`

	ID           int        `sql:"id,pk" json:"id"`
	UserID       int        `sql:"user_id,notnull" json:"user"`
	LocationID   int        `sql:"location_id,notnull" json:"-"`
	Location     Location   `json:"location"`
	StartTime    time.Time  `sql:"start_time,notnull" json:"start_time"`
	EndTime      *time.Time `sql:"end_time" json:"end_time"`
	Deleted      bool       `sql:"deleted" json:"deleted,omitempty"`
	DateCreated  time.Time  `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time  `sql:"date_modified,notnull" json:"date_modified"`
}

// BeforeInsert sets date fields
func (s *Session) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if s.DateCreated.IsZero() {
		s.DateCreated = now
	}
	s.DateModified = now
	return nil
}

// BeforeUpdate sets the date_modified field
func (s *Session) BeforeUpdate(db orm.DB) error {
	s.DateModified = time.Now()
	return nil
}

// BelongsToUser checks if the session belongs to the given user
func (s *Session) BelongsToUser(uid int, db orm.DB) bool {
	err := db.Model(s).
		WherePK().
		Where("user_id = ?", uid).
		Select()
	if err != nil {
		return false
	}
	return true
}
