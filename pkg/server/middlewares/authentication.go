package middlewares

import (
	"context"
	"cybersafe-backend-api/pkg/models"
	"net/http"
)

const UserKey = Key("user")

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//abre o jwt

		//verifica se é valido

		//pega id do user

		//busca no banco

		//add no contexto

		userData := models.User{
			Name: "Vino!",
		}

		ctx := context.WithValue(r.Context(), UserKey, userData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
