package database

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/pallavi/sendit-api/pkg/config"
)

// New initializes a new database struct
func New(cfg *config.Config) *pg.DB {
	pgOptions, err := pg.ParseURL(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(pgOptions)

	if cfg.DatabaseDebug {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
