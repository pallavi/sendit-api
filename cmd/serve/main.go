package main

import (
	"fmt"
	"log"

	"github.com/pallavi/sendit-api/pkg/application"
	"github.com/pallavi/sendit-api/pkg/server"
)

func main() {
	app := application.New()
	srv := server.New(app)

	log.Fatal(srv.Start(fmt.Sprintf(":%d", app.Config.Port)))
	log.Printf("server started on port %d", app.Config.Port)
}
