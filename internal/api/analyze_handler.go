package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func SetLogger(l *logrus.Logger) {
	logger = l
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.WithError(err).Error("Failed to parse form")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Sample Response")
}
