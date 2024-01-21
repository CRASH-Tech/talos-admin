package logger

import (
	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	log "github.com/sirupsen/logrus"
)

func Init(cfg config.СonfigImpl) {
	switch cfg.LOG_FORMAT {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	switch cfg.LOG_LEVEL {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
