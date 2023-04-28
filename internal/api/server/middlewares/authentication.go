package middlewares

import (
	"context"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/jwtutil"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const UserKey = Key("user")
const JWTIDKey = Key("jwtID")

func Authorizer(c *components.Components, allowedRoles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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

			if jwtutil.IsBlackListed(c.Cache, jwtClaims.RegisteredClaims.ID) {
				components.HttpErrorMiddlewareResponse(
					w, r,
					http.StatusUnauthorized,
					errutil.ErrInvalidJWT,
				)
				return
			}

			userID := uuid.MustParse(jwtClaims.UserID)
			user := models.User{}

			c.DB.First(&user, userID)

			if !govalidator.IsIn(jwtClaims.Role, allowedRoles...) && len(allowedRoles) != 0 {
				components.HttpErrorMiddlewareResponse(
					w, r,
					http.StatusForbidden,
					errutil.ErrInvalidJWT,
				)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, &user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}
