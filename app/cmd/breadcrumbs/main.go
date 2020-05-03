package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/tonyOreglia/breadcrumbs/internal/server"
	"github.com/tonyOreglia/breadcrumbs/config"
)

func main() {
	config := config.New()
	server := server.New(config)
	log.Info("Starting server")
	log.Fatal(server.Start())
}
