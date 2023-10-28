package users

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"fmt"
	"log"
	"os"

	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/aws"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ListUsersHandler
//
//	@Summary	List users with paginated response
//
//	@Tags		User
//	@success	200		{array}	pagination.PaginationData{data=ResponseContent}	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		page	query		int					false	"Page number"
//	@Param		limit	query		int					false	"Limit of elements per page"
//	@Router		/users [get]
//	@Security	Bearer
//	@Security	Language
func ListUsersHandler(c *components.HTTPComponents) {
	paginationData, err := pagination.GetPaginationData(c.HttpRequest.URL.Query())

	if errors.Is(err, errutil.ErrInvalidPageParam) {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	} else if errors.Is(err, errutil.ErrInvalidLimitParam) {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	users, count := c.Components.Resources.Users.ListWithPagination(paginationData.Offset, paginationData.Limit)

	response := paginationData.ToResponse(
		ToListResponse(users), int(count),
	)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// GetAuthenticatedUserHandler retrieves a user by ID
//
//	@Summary	Get authenticated user
//	@Tags		User
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Router		/users/me [get]
//	@Security	Bearer
//	@Security	Language
func GetAuthenticatedUserHandler(c *components.HTTPComponents) {
	user, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponseWithPayload(c, ToResponse(*user), http.StatusOK)
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

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	user, err := c.Components.Resources.Users.GetByID(uuid.MustParse(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
//	@Success	200		{object}	ResponseContent	"OK"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		request	body		RequestContent		true	"Request payload for creating a new user"
//	@Router		/users [post]
//	@Security	Bearer
//	@Security	Language
func CreateUserHandler(c *components.HTTPComponents) {

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUnexpectedError)
		return
	}

	userRequest := RequestContent{}
	err := components.ValidateRequest(c, &userRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := userRequest.ToEntity()

	if models.RoleToHierarchyNumber(user.Role) >= models.RoleToHierarchyNumber(currentUser.Role) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUserRole)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUnexpectedError)
		return
	}

	user.Password = string(hashedPassword)
	user.Enabled = true

	err = c.Components.Resources.Users.Create(user)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorResponse(c, http.StatusConflict, errutil.ErrCPFOrEmailAlreadyInUse)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(*user), http.StatusOK)
}

// PreSignupUserHandler
//
//	@Summary	Pre signup an user
//
//	@Tags		User
//	@Success	204		"No content"
//	@Failure	400		"Bad Request"
//	@Param		request	body	PreSignupRequest	true	"Request payload for pre signup an user"
//	@Router		/users/pre-signup [post]
//	@Security	Bearer
//	@Security	Language
func PreSignupUserHandler(c *components.HTTPComponents) {

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUnexpectedError)
		return
	}

	preSignUpRequest := PreSignupRequest{}
	err := components.ValidateRequest(c, &preSignUpRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Role:  preSignUpRequest.Role,
		Email: preSignUpRequest.Email,
	}

	if models.RoleToHierarchyNumber(user.Role) > models.RoleToHierarchyNumber(currentUser.Role) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUserRole)
		return
	}

	err = c.Components.Resources.Users.Create(user)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorResponse(c, http.StatusNotFound, errutil.ErrEmailAlreadyInUse)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(*user), http.StatusNoContent)
}

// DeleteUserHandler
//
//	@Summary	Delete a user by ID
//
//	@Tags		User
//	@Success	204		"No content"
//	@Failure	400		"Bad Request"
//	@Response	default	{object}	components.Response	"Standard error example object"
//	@Param		id		path		string				true	"ID of the user to be deleted"
//	@Router		/users/{id} [delete]
//	@Security	Bearer
//	@Security	Language
func DeleteUserHandler(c *components.HTTPComponents) {
	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	err := c.Components.Resources.Users.Delete(uuid.MustParse(id))

	if err != nil {
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

	id := chi.URLParam(c.HttpRequest, "id")

	if !govalidator.IsUUID(id) {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrInvalidUUID)
		return
	}

	profilePictureFile, err := helpers.ConvertBase64ImageToFile(userRequest.ProfilePictureURL)
	if err != nil {
		log.Println("Error converting base64 to file:", err)
	}

	profilePictureURL := fmt.Sprintf("profile-pictures/%s", profilePictureFile.Name())
	s3Client := aws.GetS3Client(aws.GetAWSConfig(c.Components))
	err = s3Client.UploadFile(c.Components.Settings.String("aws.usersBucketName"), profilePictureURL, profilePictureFile)
	if err != nil {
		log.Println("Error uploading file to S3 bucket:", err)
	}

	err = os.Remove(profilePictureFile.Name())
	if err != nil {
		log.Println("Error deleting file from local storage:", err)
	}

	user := userRequest.ToEntity()
	user.ID = uuid.MustParse(id)
	user.ProfilePictureURL = c.Components.Settings.String("aws.usersBucketURL") + profilePictureURL

	rowsAffected, err := c.Components.Resources.Users.Update(user)

	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}
	if rowsAffected == 0 {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUserResourceNotFound)
		return
	}

	response := ToResponse(*user)

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// PersonalityTestHandler
//
//	@Summary	Store personality test result
//
//	@Tags		User
//	@Success	204		"No content"
//	@Failure	400		"Bad Request"
//	@Failure	404		"User not found"
//	@Response	default	{object}	components.Response		"Standard error example object"
//	@Param		request	body		PersonalityTestRequest	true	"Request payload for personality test result"
//	@Router		/users/personality-test [post]
//	@Security	Bearer
//	@Security	Language
func PersonalityTestHandler(c *components.HTTPComponents) {
	personalityTestRequest := PersonalityTestRequest{}
	err := components.ValidateRequest(c, &personalityTestRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	currentUser, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)
	if !ok {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUnexpectedError)
		return
	}

	currentUser.MBTIType = personalityTestRequest.MBTIType

	_, err = c.Components.Resources.Users.Update(currentUser)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}
