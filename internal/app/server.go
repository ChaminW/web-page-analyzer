package app

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/chaminw/web-page-analyzer/internal/api"
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
	s.Router.HandleFunc("/analyze", api.AnalyzeHandler).Methods("POST")
}
