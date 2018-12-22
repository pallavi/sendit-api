package main

import (
	"fmt"
	"log"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/pallavi/sendit-api/pkg/application"
	"github.com/pallavi/sendit-api/pkg/server"
)

func main() {
	app := application.New()
	srv := server.New(app)

	srv.Addr = fmt.Sprintf(":%d", app.Config.Port)
	log.Printf("server started on port %d", app.Config.Port)
	gracehttp.Serve(srv)
}
