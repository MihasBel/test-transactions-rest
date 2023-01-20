package main

import (
	"flag"

	"github.com/MihasBel/test-transactions/internal/app"
	"github.com/MihasBel/test-transactions/pkg/logger"
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", "configs/local/env.json", "Config file path")
	flag.Parse()

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&app.Config, configPath); err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	log.Logger = logger.New(app.Config)
	log.Info().Msg("Application is working")
}
