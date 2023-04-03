package app

import (
	"log"
	"net"
	"net/http"

	"github.com/sinyavcev/authorization/config"
	httpController "github.com/sinyavcev/authorization/internal/controller/http"

	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	config         config.Config
	httpController *httpController.HttpController
}

func NewServer(config config.Config, httpController *httpController.HttpController) *HttpServer {
	return &HttpServer{
		config:         config,
		httpController: httpController,
	}
}

func (h *HttpServer) Run() {
	router := chi.NewRouter()
	addr := net.JoinHostPort(h.config.HttpServer.Host, h.config.HttpServer.Port)

	h.httpController.SetupAuthRoutes(router)
	srv := &http.Server{
		Addr:           addr,
		ReadTimeout:    h.config.HttpServer.ReadTimeout,
		WriteTimeout:   h.config.HttpServer.WriteTimeout,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
