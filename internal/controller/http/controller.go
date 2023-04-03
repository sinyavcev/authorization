package http

import (
	"github.com/go-chi/chi/v5"
)

type BackendUsecases interface {
	signIn()
	signUp()
	me()
	logout()
	refresh()
}

type HttpController struct {
	backendUsecases BackendUsecases
}

func NewController(backendUsecases BackendUsecases) *HttpController {
	return &HttpController{backendUsecases: backendUsecases}
}

func (h *HttpController) SetupAuthRoutes(router *chi.Mux) {
	baseURL := "/auth"
	router.Route(baseURL, func(router chi.Router) {
		router.Post("/signin", h.signUp)
		router.Get("/signup", h.signIn)
		router.Post("/refresh", h.refreshToken)
		router.Get("/me", h.me)
		router.Get("/logout", h.logout)
	})
}
