package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/chaminw/web-page-analyzer/internal/models"
	"github.com/chaminw/web-page-analyzer/internal/services"
	"github.com/chaminw/web-page-analyzer/internal/utils"
)

var (
	logger *logrus.Logger
	urlAnalyzerService *services.URLAnalyzerService
)

func SetLogger(l *logrus.Logger) {
	logger = l
}

func SetURLAnalyzerService(service *services.URLAnalyzerService) {
	urlAnalyzerService = service
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.WithError(err).Error("Failed to parse form")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	urlStr := strings.TrimSpace(r.FormValue("url"))
	if urlStr == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	if !utils.IsValidURL(urlStr) {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	result, err := urlAnalyzerService.AnalyzeURL(urlStr)
	if err != nil {
		logger.WithError(err).WithField("url", urlStr).Error("Failed to analyze URL")
		
		if httpErr, ok := err.(*models.HTTPError); ok {
			errorResponse := map[string]interface{}{
				"error":        "Failed to analyze URL",
				"status_code":  httpErr.StatusCode,
				"description":  httpErr.Description,
				"url":          urlStr,
			}
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		
		http.Error(w, "Failed to analyze URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
