package main

import (
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Panic(err)
	}

	server.Start(cfg)
}
