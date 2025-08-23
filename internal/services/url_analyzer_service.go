package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chaminw/web-page-analyzer/internal/errors"
	"github.com/chaminw/web-page-analyzer/internal/models"
	"github.com/sirupsen/logrus"
)

type URLAnalyzerService struct {
	logger *logrus.Logger
}

func NewURLAnalyzerService(l *logrus.Logger) *URLAnalyzerService {
	logger := l
	return &URLAnalyzerService{
		logger: logger,
	}
}

func (s *URLAnalyzerService) AnalyzeURL(urlStr string) (*models.AnalysisResult, error) {
	s.logger.WithField("url", urlStr).Info("Starting URL analysis")

	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create request")
		return nil, errors.Wrap("REQUEST_CREATE_ERROR", "Failed to create request", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; WebPageAnalyzer/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch URL")
		return nil, errors.Wrap("URL_FETCH_ERROR", "Failed to fetch URL", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		description := getHTTPStatusDescription(resp.StatusCode)
		s.logger.WithFields(logrus.Fields{
			"statusCode":  resp.StatusCode,
			"description": description,
		}).Warn("Failure status code received")
		return nil, &models.HTTPError{
			StatusCode:  resp.StatusCode,
			Description: description,
		}
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		s.logger.WithError(err).Error("Failed to parse HTML")
		return nil, errors.Wrap("HTML_PARSE_ERROR", "Failed to parse HTML", err)
	}

	// Analyze the document
	result := &models.AnalysisResult{
		URL:          urlStr,
		AnalysisTime: time.Now(),
	}

	// Analyze HTML version
	result.HTMLVersion = s.analyzeHTMLVersion(doc)

	// Analyze title
	result.Title = s.analyzeTitle(doc)

	// Analyze headings
	result.Headings = s.analyzeHeadings(doc)

	// Analyze links
	internal, external, inaccessible := s.analyzeLinks(doc, urlStr)
	result.InternalLinks = internal
	result.ExternalLinks = external
	result.InaccessibleLinks = inaccessible

	// Check for login form
	result.HasLoginForm = s.analyzeLoginForm(doc)

	return result, nil
}

// Determines HTML version of the document
func (s *URLAnalyzerService) analyzeHTMLVersion(doc *goquery.Document) string {
	// Check DOCTYPE declaration
	doctype := doc.Find("html").Get(0)
	if doctype != nil {
		for _, attr := range doctype.Attr {
			if attr.Key == "lang" {
				return "HTML5"
			}
		}
	}

	// Check for HTML5
	if doc.Find("header, nav, main, section, article, aside, footer").Length() > 0 {
		return "HTML5"
	}

	// Check for XHTML
	if doc.Find("html[xmlns]").Length() > 0 {
		return "XHTML"
	}

	// Default to HTML4
	return "HTML4"
}

// Extracts the page title
func (s *URLAnalyzerService) analyzeTitle(doc *goquery.Document) string {
	title := doc.Find("title").Text()
	return strings.TrimSpace(title)
}

// Counts headings by level
func (s *URLAnalyzerService) analyzeHeadings(doc *goquery.Document) map[string]int {
	headings := make(map[string]int)

	for i := 1; i <= 6; i++ {
		selector := fmt.Sprintf("h%d", i)
		count := doc.Find(selector).Length()
		if count > 0 {
			headings[fmt.Sprintf("h%d", i)] = count
		}
	}

	return headings
}

// Analyzes internal and external links
func (as *URLAnalyzerService) analyzeLinks(doc *goquery.Document, baseURL string) (internal, external, inaccessible int) {
	baseURLParsed, err := url.Parse(baseURL)
	if err != nil {
		return 0, 0, 0
	}

	var wg sync.WaitGroup
	internalChan := make(chan int, 1)
	externalChan := make(chan int, 1)
	inaccessibleChan := make(chan int, 1)

	// Find all links
	links := doc.Find("a[href]")

	wg.Add(1)
	go func() {
		defer wg.Done()
		internalCount := 0
		externalCount := 0
		inaccessibleCount := 0

		links.Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if !exists {
				return
			}

			if strings.HasPrefix(href, "javascript:") || strings.HasPrefix(href, "mailto:") {
				return
			}

			linkURL, err := url.Parse(href)
			if err != nil {
				return
			}

			if !linkURL.IsAbs() {
				linkURL = baseURLParsed.ResolveReference(linkURL)
			}

			if linkURL.Hostname() == baseURLParsed.Hostname() {
				internalCount++
			} else {
				externalCount++
				if !as.isLinkAccessible(linkURL.String()) {
					inaccessibleCount++
				}
			}
		})

		internalChan <- internalCount
		externalChan <- externalCount
		inaccessibleChan <- inaccessibleCount
	}()

	wg.Wait()

	internal = <-internalChan
	external = <-externalChan
	inaccessible = <-inaccessibleChan

	return internal, external, inaccessible
}

// Checks if a link is accessible
func (s *URLAnalyzerService) isLinkAccessible(urlStr string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Head(urlStr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// Xhecks if the page contains a login form
func (s *URLAnalyzerService) analyzeLoginForm(doc *goquery.Document) bool {
	forms := doc.Find("form")

	var hasLoginForm bool

	forms.Each(func(i int, s *goquery.Selection) {
		// Check for password field
		if s.Find("input[type='password']").Length() > 0 {
			usernameSelectors := []string{
				"input[name*='user']",
				"input[name*='email']",
				"input[name*='login']",
				"input[name*='username']",
				"input[id*='user']",
				"input[id*='email']",
				"input[id*='login']",
				"input[id*='username']",
			}

			for _, selector := range usernameSelectors {
				if s.Find(selector).Length() > 0 {
					hasLoginForm = true
					return
				}
			}
		}
	})

	return hasLoginForm
}

func getHTTPStatusDescription(statusCode int) string {
	descriptions := map[int]string{
		400: "Bad Request - The server cannot process the request due to a client error",
		401: "Unauthorized - Authentication is required to access this resource",
		403: "Forbidden - Access to this resource is denied",
		404: "Not Found - The requested resource was not found on the server",
		500: "Internal Server Error - The server encountered an unexpected condition",
		502: "Bad Gateway - The server received an invalid response from an upstream server",
		503: "Service Unavailable - The server is temporarily unable to handle the request",
		504: "Gateway Timeout - The server did not receive a timely response from an upstream server",
	}

	if desc, ok := descriptions[statusCode]; ok {
		return desc
	}
	return "An error occurred while processing the request"
}
