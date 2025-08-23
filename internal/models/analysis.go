package models

import "time"

type HTTPError struct {
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}

func (e *HTTPError) Error() string {
	return "HTTP " + string(e.StatusCode) + ": " + e.Description
}

type AnalysisResult struct {
	URL               string         `json:"url"`
	HTMLVersion       string         `json:"html_version"`
	Title             string         `json:"title"`
	Headings          map[string]int `json:"headings"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
	AnalysisTime      time.Time      `json:"analysis_time"`
}

type AnalysisRequest struct {
	URL string `json:"url" form:"url"`
}
