package user

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/models"
	"cybersafe-backend-api/pkg/pagination"
	"cybersafe-backend-api/pkg/server/middlewares"
	"net/http"
)

func ListUsersHandler(c *components.HTTPComponents) {

	dbConn := db.MustGetDbConn()

	paginationData := c.HttpRequest.Context().
		Value(middlewares.PaginationKey).(pagination.PaginationData)

	var users []models.User
	dbConn.Offset(paginationData.Offset).Limit(paginationData.Limit).Find(&users)

	var count int64
	dbConn.Model(&models.User{}).Count(&count)

	response := paginationData.ToResponse(
		ToListResponse(users), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}
