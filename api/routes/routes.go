package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jaimeiherrera/schmo_users_go/api/handler"
)

func SetupRoutes(router *chi.Mux, handlers *handler.Handlers) {
	// Global middlewares (applied to all routes)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", handlers.Ping)

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", handlers.GetUser)
			r.Get("/", handlers.GetUsers)
			r.Post("/", handlers.CreateUser)
			r.Delete("/{id}", handlers.DeleteUser)
			r.Put("/{id}", handlers.UpdateUser)
		})
	})
}
