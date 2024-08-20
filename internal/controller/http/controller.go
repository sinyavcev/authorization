package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
)

type BackendUsecases interface {
	SignUp(data *backendModels.SignUpRequest) (*backendModels.SignUpResponse, error)
	SignIn(data *backendModels.SignInRequest) (*backendModels.SignInResponse, error)
}

type Logger interface {
	Errorf(format string, args ...interface{})
}

type HttpController struct {
	backendUsecases BackendUsecases
	logger          Logger
}

func NewController(backendUsecases BackendUsecases, logger Logger) *HttpController {
	return &HttpController{
		logger: logger}
}

func (h *HttpController) SetupAuthRoutes(router *chi.Mux) {
	baseURL := "/auth"
	router.Route(baseURL, func(router chi.Router) {
		router.Post("/signin", h.SignUp)
		router.Get("/signup", h.SignIn)
	})
}
