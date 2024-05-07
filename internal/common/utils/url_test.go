package utils

import (
	"testing"
)

func TestGetDomainName(t *testing.T) {
	tests := []struct {
		inputURL string
		expected string
	}{
		{"https://www.example.com/path/to/page", "www.example.com"},
		{"https://subdomain.example.com", "subdomain.example.com"},
		{"https://example.com", "example.com"},
		{"https://example.com:8080", "example.com"},
		{"invalidurl", ""},
	}

	for _, test := range tests {
		t.Run(test.inputURL, func(t *testing.T) {
			actual, err := GetDomainName(test.inputURL)
			if err != nil {
				if test.expected != "" {
					t.Errorf("Expected domain name: %s, got error: %v", test.expected, err)
				}
				return
			}
			if actual != test.expected {
				t.Errorf("Expected domain name: %s, got: %s", test.expected, actual)
			}
		})
	}
}
