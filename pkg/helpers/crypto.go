package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func MustGenerateURLEncodedRandomToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
