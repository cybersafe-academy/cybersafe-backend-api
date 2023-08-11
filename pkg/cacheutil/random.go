package cacheutil

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	ForgotPasswordPrefix string = "forgotPassword"
	FirstAccessPrefix    string = "firstAccess"
)

func KeyWithPrefix(prefix, token string) string {
	return fmt.Sprintf("%s%s", prefix, token)
}

func MustGenRandomToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(
		[]byte(bytes))
}
