package app

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"web-page-analyzer/internal/api"
	"web-page-analyzer/internal/middleware"
)

type Server struct {
	Router *mux.Router
	logger *logrus.Logger
}

func NewServer(logger *logrus.Logger) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		logger: logger,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// Register middlewares
	s.Router.Use(middleware.LoggingMiddleware(s.logger))
	s.Router.Use(middleware.RecoveryMiddleware(s.logger))

	// Register routes
	s.Router.HandleFunc("/analyze", api.AnalyzeHandler).Methods("POST")
}
