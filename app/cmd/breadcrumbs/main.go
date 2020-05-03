package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/tonyOreglia/breadcrumbs/config"
	"github.com/tonyOreglia/breadcrumbs/internal/server"
)

func main() {
	config := config.New()
	server := server.New(config)
	log.Fatal(server.Start())
}
