package users

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"

	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/aws"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/pagination"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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
		components.HttpErrorLocalizedResponse(c, http.StatusNotFound, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInvalidUUID",
		}))
		return
	}

	user, err := c.Components.Resources.Users.GetByID(uuid.MustParse(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			components.HttpErrorLocalizedResponse(c, http.StatusNotFound, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUserResourceNotFound",
			}))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	userRequest := RequestContent{}
	err := components.ValidateRequest(c, &userRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := userRequest.ToEntity()

	err = aws.HandleImageAndUploadToS3(
		userRequest.ProfilePictureURL,
		c.Components.Settings.String("aws.usersBucketName"),
		c.Components.Settings.String("aws.usersProfilePictureFolder"),
		c.Components.Settings.String("aws.usersbucketURL"),
		c,
		user,
		1280,
		720,
	)
	if err != nil {
		errKey := "ErrUnexpectedError"

		if errors.Is(err, errutil.ErrInvalidBase64Image) {
			errKey = "ErrInvalidBase64Image"
		}

		components.HttpErrorLocalizedResponse(
			c,
			http.StatusInternalServerError,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: errKey}),
		)

		return
	}

	if models.RoleToHierarchyNumber(user.Role) > models.RoleToHierarchyNumber(currentUser.Role) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInvalidUserRole",
		}))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	user.Password = string(hashedPassword)
	user.Enabled = true

	err = c.Components.Resources.Users.Create(user)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorLocalizedResponse(c, http.StatusConflict, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrCPFOrEmailAlreadyInUse",
			}))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	preSignUpRequest := PreSignupRequest{}
	err := components.ValidateRequest(c, &preSignUpRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Role:      preSignUpRequest.Role,
		Email:     preSignUpRequest.Email,
		CompanyID: preSignUpRequest.CompanyID,
	}

	if models.RoleToHierarchyNumber(user.Role) > models.RoleToHierarchyNumber(currentUser.Role) {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInvalidUserRole",
		}))
		return
	}

	err = c.Components.Resources.Users.Create(user)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorLocalizedResponse(c, http.StatusNotFound, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrEmailAlreadyInUse",
			}))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
			return
		}
	}

	components.HttpResponseWithPayload(c, ToResponse(*user), http.StatusCreated)
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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInvalidUUID",
		}))
		return
	}

	err := c.Components.Resources.Users.Delete(uuid.MustParse(id))

	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrInvalidUUID",
		}))
		return
	}

	user := userRequest.ToEntity()
	user.ID = uuid.MustParse(id)

	err = aws.HandleImageAndUploadToS3(
		userRequest.ProfilePictureURL,
		c.Components.Settings.String("aws.usersBucketName"),
		c.Components.Settings.String("aws.usersProfilePictureFolder"),
		c.Components.Settings.String("aws.usersbucketURL"),
		c,
		user,
		1280,
		720,
	)
	if err != nil {
		errKey := "ErrUnexpectedError"

		if errors.Is(err, errutil.ErrInvalidBase64Image) {
			errKey = "ErrInvalidBase64Image"
		}

		components.HttpErrorLocalizedResponse(
			c,
			http.StatusInternalServerError,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: errKey}),
		)

		return
	}

	rowsAffected, err := c.Components.Resources.Users.Update(user)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}
	if rowsAffected == 0 {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUserResourceNotFound",
		}))
		return
	}

	compay, err := c.Components.Resources.Companies.GetByID(user.CompanyID)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	user.Company = compay

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
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	currentUser.MBTIType = personalityTestRequest.MBTIType

	_, err = c.Components.Resources.Users.Update(currentUser)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError, c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "ErrUnexpectedError",
		}))
		return
	}

	components.HttpResponse(c, http.StatusNoContent)
}
