package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TABLE users (
				id 			  serial PRIMARY KEY,
				username 	  varchar(100) UNIQUE NOT NULL,
				password 	  varchar(60) NOT NULL,
				date_created  timestamptz NOT NULL,
				date_modified timestamptz NOT NULL
			)
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("DROP TABLE users")
		return err
	})
}
