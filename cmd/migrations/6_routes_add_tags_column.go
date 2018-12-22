package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		if _, err := db.Exec("ALTER TABLE routes ADD COLUMN tags VARCHAR(100)[]"); err != nil {
			return err
		}
		_, err := db.Exec("CREATE INDEX routes_tags_idx ON routes USING GIN(tags)")
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec("ALTER TABLE routes DROP COLUMN tags")
		return err
	})
}
