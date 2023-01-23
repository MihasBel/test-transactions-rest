package main

import (
	"context"
	"flag"
	"github.com/MihasBel/test-transactions-rest/adapter/client/grpc"
	"github.com/google/uuid"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MihasBel/test-transactions-rest/internal/broker"
	"github.com/MihasBel/test-transactions-rest/internal/rep"

	"github.com/MihasBel/test-transactions-rest/delivery"
	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/MihasBel/test-transactions-rest/pkg/logger"
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
	l := logger.New(app.Config)
	g := grpc.New(app.Config.GRPC, &l)
	if err := g.Start(context.Background()); err != nil {
		l.Fatal().Msg("")
	}

	t, err := g.GetTransaction(context.Background(), uuid.MustParse("1fe1cda9-4697-48f6-a9e3-9f523894dbc6"))
	if err != nil {
		l.Fatal().Msg("")
	}
	l.Info().Interface("tx", t).Msg("tx")
	log.Logger = l
	//TODO start DB
	b := broker.NewBroker(app.Config, l)
	h := rep.NewBTransactor(b, l, app.Config)
	application := delivery.New(app.Config, h)
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
