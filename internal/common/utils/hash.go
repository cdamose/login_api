package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashURL(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashBytes)[:8]
}
