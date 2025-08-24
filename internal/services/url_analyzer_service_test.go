package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzeURL(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head><title>Test Page</title></head>
			<body>
				<h1>Header 1</h1>
				<h2>Header 2</h2>
				<a href="/internal">Internal Link</a>
				<a href="https://external.com">External Link</a>
				<form>
					<input type="text" name="username" />
					<input type="password" name="password" />
				</form>
			</body>
			</html>
		`))
	}))
	defer ts.Close()

	logger := logrus.New()
	service := NewURLAnalyzerService(logger)
	result, err := service.AnalyzeURL(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Test Page", result.Title)
	assert.Equal(t, "HTML4", result.HTMLVersion)
	assert.Equal(t, 1, result.InternalLinks)
	assert.Equal(t, 1, result.ExternalLinks)
	assert.True(t, result.HasLoginForm)
}
