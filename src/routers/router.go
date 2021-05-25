package routers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (routes *route) RouterGroup() http.Handler {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Get("/test-db", routes.ctrl.TestDB)
		r.Get("/account/{account_number}", routes.ctrl.GetBalance)
	})
	return router
}
