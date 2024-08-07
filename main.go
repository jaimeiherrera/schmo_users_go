package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jaimeiherrera/schmo_users_go/api"
	"github.com/jaimeiherrera/schmo_users_go/api/handler"
	"github.com/jaimeiherrera/schmo_users_go/api/routes"
)

func main() {
	r := chi.NewRouter()
	components := api.NewComponents()
	handlers := handler.NewHandlers(*components)
	routes.SetupRoutes(r, handlers)

	http.ListenAndServe(":8080", r)
}
