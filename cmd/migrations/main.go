package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/go-pg/migrations"
	"github.com/pallavi/sendit-api/pkg/config"
	"github.com/pallavi/sendit-api/pkg/database"
)

const directory = "cmd/migrations/"

func main() {
	cfg := config.LoadConfig()
	db := database.New(cfg)

	flag.Parse()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}
	if newVersion > oldVersion {
		log.Printf("migrated from %s to %s\n", getMigrationFilename(oldVersion), getMigrationFilename(newVersion))
	} else if newVersion < oldVersion {
		log.Printf("rolled back from %s to %s\n", getMigrationFilename(oldVersion), getMigrationFilename(newVersion))
	} else {
		log.Printf("migrations are up to date: %s\n", getMigrationFilename(oldVersion))
	}
}

func getMigrationFilename(version int64) string {
	if version == 0 {
		return "empty database"
	}
	files, err := filepath.Glob(fmt.Sprintf("%s%d_*.go", directory, version))
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatal("no migrations match this version number")
	}
	if len(files) > 1 {
		log.Fatal("multiple migrations match this version number")
	}

	filename := strings.TrimPrefix(files[0], directory)
	return filename
}
