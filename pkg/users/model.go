package users

import (
	"time"

	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
)

// User represents the database model for a user
type User struct {
	tableName struct{} `sql:"users"`

	ID           int       `sql:"id,pk" json:"id"`
	JWTToken     string    `sql:"-" json:"jwt_token,omitempty"`
	Username     string    `sql:"username,notnull" json:"username"`
	Password     string    `sql:"password,notnull" json:"-"`
	DateCreated  time.Time `sql:"date_created,notnull" json:"date_created"`
	DateModified time.Time `sql:"date_modified,notnull" json:"date_modified"`
}

// BeforeInsert hashes password and sets date fields
func (u *User) BeforeInsert(db orm.DB) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash

	now := time.Now()
	if u.DateCreated.IsZero() {
		u.DateCreated = now
	}
	u.DateModified = now
	return nil
}

// BeforeUpdate hashes password and sets the date_modified field
func (u *User) BeforeUpdate(db orm.DB) error {
	if u.Password != "" {
		hash, err := hashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hash
	}

	u.DateModified = time.Now()
	return nil
}

func hashPassword(password string) (string, error) {
	cost := 8
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
