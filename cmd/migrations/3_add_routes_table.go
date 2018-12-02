package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		if _, err := db.Exec(`
			CREATE TABLE routes (
				id 			  serial PRIMARY KEY,
				user_id 	  integer REFERENCES users (id) ON DELETE CASCADE NOT NULL,
				location_id   integer REFERENCES locations (id) ON DELETE CASCADE NOT NULL,
				name 		  varchar(100) NOT NULL,
				type 		  varchar(100) CHECK (type IN ('boulder', 'toprope', 'lead')) NOT NULL,
				grade 		  varchar(10) NOT NULL,
				deleted       boolean DEFAULT false NOT NULL,
				date_created  timestamptz NOT NULL,
				date_modified timestamptz NOT NULL
			)
		`); err != nil {
			return err
		}
		_, err := db.Exec("CREATE INDEX routes_user_id_location_id_idx ON routes (user_id, location_id)")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("DROP TABLE routes")
		return err
	})
}
