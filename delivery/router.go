package delivery

func (r *REST) setURLs() {
	api := r.app.Group("/api")
	v1 := api.Group("/v1")
	transaction := v1.Group("/transaction")
	transaction.Post("/", r.place)
	transaction.Get("/:transactionId", r.byID)

}
