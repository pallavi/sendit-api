package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		if _, err := db.Exec(`
			CREATE TABLE sessions (
				id 			  serial PRIMARY KEY,
				user_id       integer REFERENCES users (id) ON DELETE CASCADE NOT NULL,
				location_id   integer REFERENCES locations (id) ON DELETE CASCADE NOT NULL,
				start_time 	  timestamptz NOT NULL,
				end_time	  timestamptz,
				deleted       boolean DEFAULT false NOT NULL,
				date_created  timestamptz NOT NULL,
				date_modified timestamptz NOT NULL
			)
		`); err != nil {
			return err
		}
		_, err := db.Exec("CREATE INDEX sessions_user_id_location_id_idx ON sessions (user_id, location_id)")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("DROP TABLE sessions")
		return err
	})
}
