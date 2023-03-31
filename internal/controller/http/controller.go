package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/internal/backendUsecases"
)

type Controller struct {
	BackendUsecases *backendUsecases.Authorization
}

func NewController(backendUsecases *backendUsecases.Authorization) *Controller {
	return &Controller{BackendUsecases: backendUsecases}
}

func (c *Controller) SetupAuthRoutes(router *chi.Mux) *chi.Mux {
	baseURL := "/auth"
	router.Route(baseURL, func(router chi.Router) {
		router.Post("/signin", c.signUp)
		router.Get("/signup", c.signIn)
		router.Post("/refresh", c.refreshToken)
		router.Get("/me", c.me)
		router.Get("/logout", c.logout)
	})

	return router
}
