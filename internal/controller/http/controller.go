package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
)

type BackendUsecases interface {
	SignIn(data backendModels.SignInRequest) (*backendModels.SignInResponse, error)
	SignUp(data backendModels.SignUpRequest) (*backendModels.SignInResponse, error)
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
	})
}
