package http

import (
	"github.com/go-chi/chi/v5"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) InitRoutes(router *chi.Mux) *chi.Mux {
	router.Route("/auth", func(router chi.Router) {
		router.Post("/register", c.register)
		router.Get("/login", c.login)
		router.Post("/refresh", c.refreshToken)
		router.Get("/me", c.me)
		router.Get("/logout", c.logout)
	})

	return router
}
