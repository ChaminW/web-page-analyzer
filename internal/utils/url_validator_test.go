package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "HTTPS URL",
			url:      "https://example.com",
			expected: true,
		},
		{
			name:     "HTTP URL",
			url:      "http://example.com",
			expected: true,
		},
		{
			name:     "URL with path",
			url:      "https://example.com/sample/path/",
			expected: true,
		},
		{
			name:     "URL with query params",
			url:      "https://example.com/query?q1=test&q2=1",
			expected: true,
		},
		{
			name:     "Subdomain",
			url:      "https://test.example.com",
			expected: true,
		},
		{
			name:     "Invalid URL - missing protocol",
			url:      "example.com",
			expected: false,
		},
		{
			name:     "Invalid URL - wrong protocol",
			url:      "ftp://example.com",
			expected: false,
		},
		{
			name:     "Invalid URL - invalid form",
			url:      "https://",
			expected: false,
		},
		{
			name:     "empty string",
			url:      "",
			expected: false,
		},
		{
			name:     "special characters",
			url:      "https://example.com/<script>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidURL(tt.url)
			assert.Equal(t, tt.expected, result, "URL: %s", tt.url)
		})
	}
}