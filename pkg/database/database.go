package database

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/pallavi/sendit-api/pkg/config"
)

// New initializes a new database struct
func New(cfg *config.Config) *pg.DB {
	addr := fmt.Sprintf("%s:%d", cfg.DatabaseHost, cfg.DatabasePort)
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Database: cfg.DatabaseName,
	})
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
