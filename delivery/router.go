package delivery

import (
	jwtware "github.com/gofiber/jwt/v3"
)

func (r *REST) setURLs() {
	api := r.app.Group("/api")
	// authorized handlers
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(r.cfg.APIKey), //TODO secret
	}))
	v1 := api.Group("/v1")
	transaction := v1.Group("/transaction")
	transaction.Post("/", r.place)
	transaction.Get("/:transactionId", r.byID)

}
