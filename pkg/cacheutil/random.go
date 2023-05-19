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

func MustGenRandomToken(prefix string) string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s%s", prefix, bytes)))
}
