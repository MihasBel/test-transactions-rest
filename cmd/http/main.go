package main

import (
	"context"
	"flag"
	"github.com/golang/mock/gomock"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MihasBel/test-transactions-rest/delivery"
	"github.com/MihasBel/test-transactions-rest/mocks"

	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/MihasBel/test-transactions-rest/pkg/logger"
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
)

var configPath string

// @title test transactions API
// @version 1.0
// @description REST API service to manage transactions
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.StringVar(&configPath, "config", "configs/local/env.json", "Config file path")
	flag.Parse()

	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&app.Config, configPath); err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	log.Logger = logger.New(app.Config)
	//TODO start DB
	application := delivery.New(app.Config, mocks.NewMockHandler(gomock.NewController(mocks.Reporter{})),
		mocks.NewMockSecurityHandler(gomock.NewController(mocks.Reporter{})))
	startCtx, startCancel := context.WithTimeout(context.Background(), time.Duration(app.Config.StartTimeout)*time.Second)
	defer startCancel()
	if err := application.Start(startCtx); err != nil {
		log.Fatal().Err(err).Msg("cannot start application") // nolint
	}

	log.Info().Msg("application started")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Duration(app.Config.StartTimeout)*time.Second)
	defer stopCancel()

	if err := application.Stop(stopCtx); err != nil {
		log.Error().Err(err).Msg("cannot stop application")
	}
	log.Info().Msg("service is down")
}
