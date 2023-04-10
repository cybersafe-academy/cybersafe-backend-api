package jwtutil

import (
	"cybersafe-backend-api/pkg/errutil"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Parse(token *jwt.Token) (any, error) {
	jwtToken, ok := token.Method.(*jwt.SigningMethodECDSA)

	if !ok {
		return nil, errutil.ErrInvalidJWT
	}

	return jwtToken, nil

}

func Generate(userID, issuer, subject, secretKey string) (string, error) {

	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    userID,
			Subject:   subject,
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	tokenString, err := token.SignedString(secretKey)
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
