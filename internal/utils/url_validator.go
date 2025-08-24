package utils

import (
	"net/url"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(?:/[\w\-\./?%&=]*)?$`)

func IsValidURL(urlStr string) bool {
	if !urlRegex.MatchString(urlStr) {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	return u.Scheme != "" && u.Host != ""
}
