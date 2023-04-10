package middlewares

import (
	"context"
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/models"
	"errors"
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
		}

		token, err := jwt.Parse(authorizationHeader, jwtutil.Parse)

		if errors.Is(err, jwt.ErrTokenMalformed) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidJWT,
			)
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidSignature,
			)
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrTokenExpiredOrPending,
			)
		} else {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidJWT,
			)
		}

		if !token.Valid {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidJWT,
			)
		}

		claims, err := jwtutil.GetClaims(*token)

		if err != nil {
			components.HttpErrorMiddlewareResponse(
				w, r,
				http.StatusUnauthorized,
				errutil.ErrInvalidClaims,
			)
		}

		userID := claims.UserID
		user := models.User{}
		db.MustGetDbConn().First(&user, userID)

		ctx := context.WithValue(r.Context(), UserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
