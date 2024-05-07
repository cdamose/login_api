package utils

import (
	"net/url"
)

func GetDomainName(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil
}
