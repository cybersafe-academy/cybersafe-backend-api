package middlewares

import (
	"context"
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/models"
	"cybersafe-backend-api/pkg/settings"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

const UserKey = Key("user")

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrCredentialsMissing,
			)
			return
		}

		token, err := jwt.ParseWithClaims(authorizationHeader, &jwtutil.CustomClaims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(settings.ExportedSettings.String("jwt.secretKey")), nil
		})

		if errors.Is(err, jwt.ErrTokenMalformed) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidJWT,
			)
			return
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidSignature,
			)
			return
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrTokenExpiredOrPending,
			)
			return
		}

		if !token.Valid {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidJWT,
			)
			return
		}

		claims, err := jwtutil.GetClaims(*token)

		if err != nil {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidClaims,
			)
			return
		}

		userID := claims.UserID
		user := models.User{}
		db.MustGetDbConn().First(&user, userID)

		ctx := context.WithValue(r.Context(), UserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
