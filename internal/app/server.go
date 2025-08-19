package app

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/chaminw/web-page-analyzer/internal/api"
	"github.com/chaminw/web-page-analyzer/internal/middleware"
	"github.com/chaminw/web-page-analyzer/internal/services"
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

	s.setupServices()
	s.setupRoutes()
	return s
}

func (s *Server) setupServices() {
	// Create service instances
	urlAnalyzerService := services.NewURLAnalyzerService()
	
	// Wire up services to API handlers
	api.SetLogger(s.logger)
	api.SetURLAnalyzerService(urlAnalyzerService)
}

func (s *Server) setupRoutes() {
	// Register middlewares
	s.Router.Use(middleware.LoggingMiddleware(s.logger))
	s.Router.Use(middleware.RecoveryMiddleware(s.logger))

	// Register routes
	s.Router.HandleFunc("/analyze", api.AnalyzeHandler).Methods("POST")
}
