package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/MihasBel/test-transactions-rest/internal/rep"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// REST represent REST-full application
type REST struct {
	app *fiber.App
	cfg Config
	t   rep.Transactor
}

// New Create new instance of REST.
func New(config Config, t rep.Transactor) *REST {
	a := fiber.New()
	a.Use(cors.New())
	rest := REST{
		app: a,
		cfg: config,
		t:   t,
	}
	rest.setURLs()

	return &rest
}

// Start rest server
func (r *REST) Start(_ context.Context) error {
	errCh := make(chan error)
	log.Debug().Msgf("start listening %q", r.cfg.Address)
	go func() {
		if err := r.app.Listen(r.cfg.Address); err != nil && err != http.ErrServerClosed {
			errCh <- errors.Wrap(err, "cannot listen and serve")
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(r.cfg.StartTimeout) * time.Second):
		return nil
	}
}

// Stop rest server
func (r *REST) Stop(_ context.Context) error {
	errCh := make(chan error)
	log.Debug().Msgf("stopping %q", r.cfg.Address)
	go func() {
		if err := r.app.Shutdown(); err != nil {
			errCh <- errors.Wrap(err, "cannot shutdown")
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(r.cfg.StopTimeout) * time.Second):
		return nil

	}
}
