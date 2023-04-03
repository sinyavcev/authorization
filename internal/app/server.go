package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/config"
	httpController "github.com/sinyavcev/authorization/internal/controller/http"
	"log"
	"net"
	"net/http"
)

type Server struct {
	config         config.Config
	httpController *httpController.Controller
}

func NewServer(config config.Config, httpController *httpController.Controller) *Server {
	return &Server{
		config:         config,
		httpController: httpController,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()
	addr := net.JoinHostPort(s.config.HttpServer.Host, s.config.HttpServer.Port)

	s.httpController.SetupAuthRoutes(router)
	srv := &http.Server{
		Addr:           addr,
		ReadTimeout:    s.config.HttpServer.ReadTimeout,
		WriteTimeout:   s.config.HttpServer.ReadTimeout,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
