package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/chaminw/web-page-analyzer/internal/api"
	"github.com/chaminw/web-page-analyzer/internal/middleware"
	"github.com/chaminw/web-page-analyzer/internal/services"
)

type Server struct {
	Router   *mux.Router
	logger   *logrus.Logger
	template *template.Template
}

func NewServer(logger *logrus.Logger) *Server {
	s := &Server{
		Router:   mux.NewRouter(),
		logger:   logger,
		template: template.Must(template.ParseGlob("web/templates/*.html")),
	}

	s.setupServices()
	s.setupRoutes()
	return s
}

func (s *Server) setupServices() {
	// Create service instances
	urlAnalyzerService := services.NewURLAnalyzerService(s.logger)

	// Wire up services to API handlers
	api.SetLogger(s.logger)
	api.SetURLAnalyzerService(urlAnalyzerService)
}

func (s *Server) setupRoutes() {
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Register routes
	s.Router.HandleFunc("/", s.HomeHandler).Methods("GET")
	s.Router.HandleFunc("/analyze", api.AnalyzeHandler).Methods("POST")

	// Register middlewares
	s.Router.Use(middleware.LoggingMiddleware(s.logger))
	s.Router.Use(middleware.RecoveryMiddleware(s.logger))
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.template.ExecuteTemplate(w, "index.html", nil); err != nil {
		s.logger.WithError(err).Error("Failed to execute template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
