package jwtutil

import (
	"cybersafe-backend-api/pkg/errutil"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Generate(userID, issuer, subject, secretKey string, expirationTime time.Duration) (string, error) {

	claims := CustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    userID,
			Subject:   subject,
			ID:        uuid.NewString(),
		},
	}

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
