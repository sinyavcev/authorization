package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/internal/usecases"
)

type HttpController struct {
	backendUsecases *usecases.Authorization
}

func NewController(backendUsecases *usecases.Authorization) *HttpController {
	return &HttpController{backendUsecases: backendUsecases}
}

func (c *HttpController) SetupAuthRoutes(router *chi.Mux) {
	baseURL := "/auth"
	router.Route(baseURL, func(router chi.Router) {
		router.Post("/signin", c.signUp)
		router.Get("/signup", c.signIn)
		router.Post("/refresh", c.refreshToken)
		router.Get("/me", c.me)
		router.Get("/logout", c.logout)
	})
}
