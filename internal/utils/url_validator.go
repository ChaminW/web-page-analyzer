package utils

import (
	"net/url"
	"regexp"
	"strings"
)

// Validates if the given string is a valid URL
func IsValidURL(urlStr string) bool {
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if u.Hostname() == "" {
		return false
	}

	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(urlStr)
}
