package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		if _, err := db.Exec(`
			CREATE TABLE locations (
				id			  serial PRIMARY KEY,
				user_id 	  integer REFERENCES users (id) ON DELETE CASCADE NOT NULL,
				name 		  varchar(100) NOT NULL,
				outdoor 	  boolean NOT NULL,
				deleted       boolean DEFAULT false NOT NULL,
				date_created  timestamptz NOT NULL,
				date_modified timestamptz NOT NULL
			)
		`); err != nil {
			return err
		}
		_, err := db.Exec("CREATE INDEX locations_user_id_idx ON locations (user_id)")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("DROP TABLE locations")
		return err
	})
}
