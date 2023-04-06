package http

import (
	"github.com/sinyavcev/authorization/internal/common/logger"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"

	"github.com/go-chi/chi/v5"
)

type BackendUsecases interface {
	SignIn(data backendModels.SignInRequest) (*backendModels.SignInResponse, error)
	SignUp(data backendModels.SignUpRequest) (*backendModels.SignInResponse, error)
}

type HttpController struct {
	backendUsecases BackendUsecases
	logger          logger.Logger
}

func NewController(backendUsecases BackendUsecases, logger *logger.Logger) *HttpController {
	return &HttpController{
		backendUsecases: backendUsecases,
		logger:          *logger}
}

func (h *HttpController) SetupAuthRoutes(router *chi.Mux) {
	baseURL := "/auth"
	router.Route(baseURL, func(router chi.Router) {
		router.Post("/signin", h.signUp)
		router.Get("/signup", h.signIn)
	})
}
