package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/config"
	authController "github.com/sinyavcev/authorization/internal/controller/http"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	config         config.Config
	authController authController.Controller
}

func NewServer(config config.Config, authController authController.Controller) *Server {
	return &Server{
		config:         config,
		authController: authController,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()
	srv := &http.Server{
		Addr:           net.JoinHostPort(s.config.HttpServer.Host, s.config.HttpServer.Port), //
		Handler:        s.authController.InitRoutes(router),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
