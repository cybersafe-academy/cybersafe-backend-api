package jwtutil

import (
	"cybersafe-backend-api/pkg/cacheutil"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/settings"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate(claims CustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetClaims(token jwt.Token) (*CustomClaims, error) {
	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, errutil.ErrInvalidClaims
	} else {
		return claims, nil
	}

}

func Parse(token string, claims *CustomClaims) (*jwt.Token, error) {
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(settings.ExportedSettings.String("jwt.secretKey")), nil
	})

	return jwtToken, err
}

func IsBlackListed(c cacheutil.Cacher, jwtID string) bool {
	_, found := c.Get(jwtID)

	return found
}

func AddToBlackList(c cacheutil.Cacher, duration time.Duration, jwtID, tokenString string) {
	c.Set(
		jwtID,
		tokenString,
		duration,
	)
}
