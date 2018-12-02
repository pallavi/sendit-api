package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		if _, err := db.Exec(`
			CREATE TABLE climbs (
				id 			  serial PRIMARY KEY,
				session_id 	  integer REFERENCES sessions (id) ON DELETE CASCADE NOT NULL,
				route_id 	  integer REFERENCES routes (id) ON DELETE CASCADE NOT NULL,
				attempts 	  integer CHECK (attempts > 0) NOT NULL,
				sent 		  boolean NOT NULL,
				deleted       boolean DEFAULT false NOT NULL,
				date_created  timestamptz NOT NULL,
				date_modified timestamptz NOT NULL
			)
		`); err != nil {
			return err
		}
		if _, err := db.Exec("CREATE INDEX climbs_session_id_idx ON climbs (session_id)"); err != nil {
			return err
		}
		_, err := db.Exec("CREATE INDEX climbs_route_id_idx ON climbs (route_id)")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("DROP TABLE climbs")
		return err
	})
}
