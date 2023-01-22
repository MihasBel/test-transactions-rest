package delivery

import (
	"context"
	"github.com/MihasBel/test-transactions-rest/internal/gen"
	"net/http"
	"time"

	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// REST represent REST-full application
type REST struct {
	cfg        app.Configuration
	h          gen.Handler
	sh         gen.SecurityHandler
	httpServer http.Server
}

// New Create new instance of REST. Should use only in main.
func New(config app.Configuration, h gen.Handler, sh gen.SecurityHandler) *REST {
	a := fiber.New()
	a.Use(cors.New())
	rest := REST{
		cfg: config,
		h:   h,
		sh:  sh,
	}
	return &rest
}

// Start an application
func (r *REST) Start(_ context.Context) error {
	log.Debug().Msgf("start listening %q", r.cfg.Address)
	errCh := make(chan error)
	oasServer, err := gen.NewServer(r.h, r.sh)
	if err != nil {
		return errors.Wrap(err, "server init")
	}
	r.httpServer = http.Server{
		Addr:    r.cfg.Address,
		Handler: oasServer,
	}
	go func() {
		if err := r.httpServer.ListenAndServe(); err != nil {
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

// Stop an application
func (r *REST) Stop(ctx context.Context) error {
	log.Debug().Msgf("stopping %q", r.cfg.Address)
	errCh := make(chan error)
	go func() {
		if err := r.httpServer.Shutdown(ctx); err != nil {
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
