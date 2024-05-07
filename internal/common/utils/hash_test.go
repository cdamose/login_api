package utils

import (
	"testing"
)

func TestHashURL(t *testing.T) {
	tests := []struct {
		inputURL string
		expected string
	}{
		{"https://example.com", "EAaArVRs"},
		{"https://google.com", "BQRvJsg-"},
		{"https://stackoverflow.com", "b2Dg_r9S"},
	}

	for _, test := range tests {
		t.Run(test.inputURL, func(t *testing.T) {
			actual := HashURL(test.inputURL)
			if actual != test.expected {
				t.Errorf("Expected hash: %s, got: %s", test.expected, actual)
			}
		})
	}
}
