package application

import (
	"github.com/go-pg/pg"
	"github.com/pallavi/sendit-api/pkg/config"
	"github.com/pallavi/sendit-api/pkg/database"
)

// App contains necessary references persisted throughout the app lifecycle
type App struct {
	Config *config.Config
	DB     *pg.DB
}

// New creates a new instance of App
func New() *App {
	cfg := config.LoadConfig()
	db := database.New(cfg)
	return &App{cfg, db}
}
