package app

import (
	"context"

	"github.com/MihasBel/test-transactions-rest/adapter/broker"
	"github.com/MihasBel/test-transactions-rest/adapter/client/grpc"
	"github.com/MihasBel/test-transactions-rest/delivery"
	"github.com/MihasBel/test-transactions-rest/internal/rep"
	"github.com/MihasBel/test-transactions-rest/pkg/cache"

	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Lifecycle to start and stop modules
type Lifecycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
type cmp struct {
	Service Lifecycle
	Name    string
}

// App represents application
type App struct {
	log  *zerolog.Logger
	cmps []cmp
	cfg  Configuration
}

// New create new app instance
func New(cfg Configuration) *App {
	l := zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().
		Str("cmp", "app").Logger()
	return &App{
		log:  &l,
		cfg:  cfg,
		cmps: []cmp{},
	}
}

// Start an application
func (a *App) Start(ctx context.Context) error {
	a.log.Info().Msg("starting app")

	redis := cache.New(a.cfg.Redis)
	g := grpc.New(a.cfg.GRPC, *a.log)
	b := broker.New(a.cfg.Kafka, *a.log)
	transactor := rep.NewBTransactor(b, g, redis, *a.log)
	rest := delivery.New(a.cfg.REST, transactor)
	a.cmps = append(
		a.cmps,
		cmp{g, "grpc"},
		cmp{rest, "rest"},
	)

	okCh, errCh := make(chan struct{}), make(chan error)
	go func() {
		for _, c := range a.cmps {
			a.log.Info().Msgf("%v is starting", c.Name)
			if err := c.Service.Start(ctx); err != nil {
				a.log.Error().Err(err).Msgf("Cannot start %v", c.Name)
				errCh <- errors.Wrapf(err, "Cannot start %v", c.Name)
			}
		}
		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("Start timeout")
	case err := <-errCh:
		return err
	case <-okCh:
		return nil
	}
}

// Stop an application
func (a *App) Stop(ctx context.Context) error {
	a.log.Info().Msg("shutting down service...")

	okCh, errCh := make(chan struct{}), make(chan error)
	go func() {
		for i := len(a.cmps) - 1; i >= 0; i-- {
			c := a.cmps[i]
			a.log.Info().Msgf("%v is stopping", c.Name)
			if err := c.Service.Start(ctx); err != nil {
				a.log.Error().Err(err).Msgf("Cannot stop %v", c.Name)
				errCh <- errors.Wrapf(err, "Cannot stop %v", c.Name)
			}
		}
		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("Stop timeout")
	case err := <-errCh:
		return err
	case <-okCh:
		return nil
	}
}
