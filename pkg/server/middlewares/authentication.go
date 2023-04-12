package middlewares

import (
	"context"
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/components/db"
	"cybersafe-backend-api/pkg/components/jwtutil"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/models"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const UserKey = Key("user")
const JWTIDKey = Key("jwtID")

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

		jwtClaims := &jwtutil.CustomClaims{}

		token, err := jwtutil.Parse(authorizationHeader, jwtClaims)

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

		userID := uuid.MustParse(jwtClaims.UserID)
		user := models.User{}

		db.MustGetDbConn().First(&user, userID)

		ctx := context.WithValue(r.Context(), UserKey, &user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
