package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/sinyavcev/authorization/config"
	"log"
	"net"
	"net/http"
)

type Server struct {
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()
	srv := &http.Server{
		Addr:           net.JoinHostPort(s.config.HttpServer.Host, s.config.HttpServer.Port),
		ReadTimeout:    s.config.HttpServer.ReadTimeout,
		WriteTimeout:   s.config.HttpServer.ReadTimeout,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
