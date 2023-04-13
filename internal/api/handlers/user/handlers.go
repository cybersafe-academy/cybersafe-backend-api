package user

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/jwtutil"
	"time"

	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ListUsersHandler
//
//	@Summary	List users with paginated response
//
//	@Tags		User
//	@success	200		{array}	pagination.PaginationData{content=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/users [get]
//	@Security	Bearer
//	@Security	Language
func ListUsersHandler(c *components.HTTPComponents) {

	dbConn := db.MustGetDbConn()

	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusNotFound, err)
		return
	}

	var users []models.User
	dbConn.Offset(paginationData.Offset).Limit(paginationData.Limit).Find(&users)

	var count int64
	dbConn.Model(&models.User{}).Count(&count)

	response := paginationData.ToResponse(
		ToListResponse(users), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// GetUserByIDHandler retrieves a user by ID
//
//	@Summary	Get user by ID
//	@Tags		User
//	@Param		id		path		string			true	"ID of the user to be retrieved"
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Router		/users/{id} [get]
//	@Security	Bearer
//	@Security	Language
func GetUserByIDHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	dbConn := db.MustGetDbConn()

	var user models.User
	result := dbConn.First(&user, uuid.MustParse(id))

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrUserResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(user), http.StatusOK)
}

// CreateUserHandler
//
//	@Summary	Create a user
//
//	@Tags		User
//	@success	200		{array}	pagination.PaginationData{content=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for creating a new user"
//	@Router		/users [post]
//	@Security	Bearer
//	@Security	Language
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
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*user), http.StatusOK)
}

// DeleteUserHandler
//
//	@Summary	Delete a user by ID
//
//	@Tags		User
//	@success	200		{array}	pagination.PaginationData{content=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of the user to be deleted"
//	@Router		/users/{id} [delete]
//	@Security	Bearer
//	@Security	Language
func DeleteUserHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	dbConn := db.MustGetDbConn()

	result := dbConn.Delete(&models.User{}, uuid.MustParse(id))

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}

// UpdateUserHandler
//
//	@Summary	Update user by ID
//
//	@Tags		User
//	@success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Failure	404		"User not found"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for updating user information"
//	@Param		id		path		string				true	"ID of user to be updated"
//	@Router		/users/{id} [put]
//	@Security	Bearer
//	@Security	Language
func UpdateUserHandler(c *components.HTTPComponents) {
	userRequest := RequestContent{}
	err := components.ValidateRequest(c, &userRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	dbConn := db.MustGetDbConn()
	id := chi.URLParam(c.HttpRequest, "id")

	user := &models.User{}
	result := dbConn.First(user, id)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrUserResourceNotFound)
		return
	}

	updatedUser := userRequest.ToEntity()
	result = dbConn.Model(user).Updates(updatedUser)

	if result.Error != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*updatedUser), http.StatusOK)
}

// LogOffHandler is the HTTP handler for user log off
//
//	@Summary		User logoff
//	@Description	Logs off an user
//	@Tags			User
//	@Success		204
//	@Failure		400	"Bad Request"
//
//	@Router			/users/logoff [post]
//	@Security		Bearer
//	@Security		Language
func LogOffHandler(c *components.HTTPComponents) {
	authorizationHeader := c.HttpRequest.Header.Get("Authorization")

	jwtClaims := &jwtutil.CustomClaims{}

	// This error cannot occur because the token was already parsed in the middleware
	token, _ := jwtutil.Parse(authorizationHeader, jwtClaims)

	jwtutil.AddToBlackList(
		time.Until(jwtClaims.RegisteredClaims.ExpiresAt.Time),
		jwtClaims.RegisteredClaims.ID,
		token.Raw,
	)

	components.HttpResponse(c, http.StatusNoContent)
}
