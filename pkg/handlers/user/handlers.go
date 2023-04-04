package user

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/errors"
	"cybersafe-backend-api/pkg/models"
	"cybersafe-backend-api/pkg/pagination"
	"cybersafe-backend-api/pkg/server/middlewares"
	"net/http"
)

// ListDomainsHandler @Summary   List domains
//
//	@Tags		Domain
//	@success	200		{array}	pagination.PaginationData{content=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/base-paths [get]
//	@Security	Bearer
//	@Security	Language
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

func CreateUserHandler(c *components.HTTPComponents) {
	userRequest := RequestContent{}
	err := components.ValidateRequest(c, &userRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := userRequest.ToEntity()
	dbConn := db.MustGetDbConn()

	result := dbConn.Create(user)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, errors.ErrUnexpectedError)
		return
	}

	response := ResponseContent{
		UserFields: userRequest.UserFields,
		ID:         user.ID,
	}

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}
