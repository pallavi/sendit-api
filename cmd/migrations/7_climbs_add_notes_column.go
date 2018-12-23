package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec("ALTER TABLE climbs ADD COLUMN notes TEXT")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("ALTER TABLE climbs DROP COLUMN notes")
		return err
	})
}
