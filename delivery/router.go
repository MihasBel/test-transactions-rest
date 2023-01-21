package delivery

func (r *REST) setURLs() {
	api := r.app.Group("/api")

	v1 := api.Group("/v1", r.isAuth)
	v1.Post("/", r.create)
	v1.Get("/history/:userId", r.history)
	v1.Get("/balance/:userId", r.balance)
	v1.Get("/byID/:transactionId", r.byID)
	v1.Get("/byTime/", r.byTime)

}
