package middlewares

import (
	"context"
	"cybersafe-backend-api/pkg/pagination"
	"net/http"
	"strconv"
)

const PaginationKey = Key("pagination")

func PaginationParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		page := 1
		limit := 10

		if pageParam := r.URL.Query().Get("page"); pageParam != "" {
			page, _ = strconv.Atoi(pageParam)

			//@TODO: Error handling
		}

		if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
			limit, _ = strconv.Atoi(limitParam)

			//@TODO: Error handling
		}

		paginationData := pagination.PaginationData{
			Limit:  limit,
			Page:   page,
			Offset: (page - 1) * limit,
		}

		ctx := context.WithValue(r.Context(), PaginationKey, paginationData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
