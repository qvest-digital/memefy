package main

import (
	"context"
	"os"

	"memefy/server/pkg/config"
	"memefy/server/pkg/server"

	log "github.com/Sirupsen/logrus"
)

const banner = `
                           ___       
  __ _  ___   __ _  ___   / _/ __ __
 /  ' \/ -_) /  ' \/ -_) / _/ / // /
/_/_/_/\__/ /_/_/_/\__/ /_/   \_, / 
                             /___/  
`

var cfg *config.Config

func init() {
	cfg = config.Get()

	// setup logrus global exported logger
	logger := log.StandardLogger()
	l, err := log.ParseLevel(cfg.Log.Level)
	if err != nil {
		log.Warnf("Failed to parse LOG_LEVEL=%s: %v", cfg.Log, err)
		l = log.InfoLevel
	}
	logger.SetLevel(l)
}

func main() {
	log.Info(banner)
	log.Infof("Config: %+v", cfg)

	err := os.MkdirAll(cfg.StoragePath, 0777)
	if err != nil {
		log.Fatal("storage not available", err)
	}

	server.RunServer(context.Background(), nil, cfg)
}
