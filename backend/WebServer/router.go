package webserver

import (
	"github.com/Nikolat27/king_of_numbers/backend/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	CoreRouter *chi.Mux
}

// newRouter -> Construction
func newRouter(handler *handlers.Handler) *Router {
	routerInstance := chi.NewRouter()
	routerInstance.Use(CheckCorsOrigin)
	routerInstance.Use(middleware.Logger)

	// All APIs should be used via 'api' prefix
	routerInstance.Route("/api", func(r chi.Router) {
		getHealthRoute(r, handler)
		getUsersRoutes(r, handler)
	})

	return &Router{CoreRouter: routerInstance}
}

func getHealthRoute(r chi.Router, handler *handlers.Handler) {
	r.Get("/health", handler.HealthCheck)
}

func getUsersRoutes(r chi.Router, handler *handlers.Handler) {
	r.Post("/users/register", handler.RegisterUser)
}
