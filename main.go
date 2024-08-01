package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jaimeiherrera/schmo_users_go/api"
)

func main() {
	r := chi.NewRouter()
	components := api.NewComponents()
	handlers := api.NewHandlers(*components)
	api.SetupRoutes(r, handlers)

	http.ListenAndServe(":3000", r)
}
