package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/aws"
	"cybersafe-backend-api/pkg/cacheutil"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/mail"
	"errors"
	"fmt"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginHandler is the HTTP handler for user login
//
//	@Summary		User login
//	@Description	Authenticates an user and generates an access token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"User login information"
//	@Success		200		{object}	TokenContent
//	@Failure		400		"Bad Request"
//
//	@Router			/auth/login [post]
func LoginHandler(c *components.HTTPComponents) {
	loginRequest := LoginRequest{}
	err := components.ValidateRequest(c, &loginRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := c.Components.Resources.Users.GetByCPF(loginRequest.CPF)

	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusUnauthorized,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUserResourceNotFound",
			}))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusUnauthorized,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrLoginOrPasswordIncorrect",
			}))
		return
	}

	expirationTime := 1 * time.Hour
	JWTID := uuid.NewString()

	claims := jwtutil.CustomClaims{
		UserID: user.ID.String(),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    c.Components.Settings.String("application.issuer"),
			Subject:   c.Components.Settings.String("jwt.subject"),
			ID:        JWTID,
		},
	}

	tokenString, err := jwtutil.Generate(
		claims,
		c.Components.Settings.String("jwt.secretKey"),
	)

	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
		return
	}

	response := TokenContent{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expirationTime.Seconds(),
	}

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// LoginHandler is the HTTP handler for user login
//
//	@Summary		User login refresh
//	@Description	Refreshes the token for an authenticated user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	TokenContent
//	@Failure		400	"Bad Request"
//
//	@Router			/auth/refresh [post]
//	@Security		Bearer
//	@Security		Language
func RefreshTokenHandler(c *components.HTTPComponents) {

	user, ok := c.HttpRequest.Context().Value(middlewares.UserKey).(*models.User)

	if !ok {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
		return
	}

	expirationTime := 1 * time.Hour
	JWTID := uuid.NewString()

	claims := jwtutil.CustomClaims{

		UserID: user.ID.String(),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    c.Components.Settings.String("application.issuer"),
			Subject:   c.Components.Settings.String("jwt.subject"),
			ID:        JWTID,
		},
	}

	tokenString, err := jwtutil.Generate(
		claims,
		c.Components.Settings.String("jwt.secretKey"),
	)

	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
		return
	}

	response := TokenContent{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expirationTime.Seconds(),
	}

	components.HttpResponseWithPayload(c, response, http.StatusOK)
}

// LogOffHandler is the HTTP handler for user log off
//
//	@Summary		User logoff
//	@Description	Logs off an user
//	@Tags			Authentication
//	@Success		204
//	@Failure		400	"Bad Request"
//
//	@Router			/auth/logoff [post]
//	@Security		Bearer
//	@Security		Language
func LogOffHandler(c *components.HTTPComponents) {
	authorizationHeader := c.HttpRequest.Header.Get("Authorization")

	jwtClaims := &jwtutil.CustomClaims{}

	// This error cannot occur because the token was already parsed in the middleware
	secretKey := c.Components.Settings.String("jwt.secretKey")
	token, _ := jwtutil.Parse(authorizationHeader, jwtClaims, secretKey)

	jwtutil.AddToBlackList(
		&c.Components.Cache,
		time.Until(jwtClaims.RegisteredClaims.ExpiresAt.Time),
		jwtClaims.RegisteredClaims.ID,
		token.Raw,
	)

	components.HttpResponse(c, http.StatusNoContent)
}

// ForgotPasswordHandler is the HTTP handler for requesting a new password
//
//	@Summary		Request new password via e-mail
//	@Description	Receives the user email and if the email is valid, send a verification via email
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	ForgotPasswordRequest	true	"Reset password info"
//
//	@Success		204		"No content"
//	@Failure		400		"Bad Request"
//
//	@Router			/auth/forgot-password [post]
func ForgotPasswordHandler(c *components.HTTPComponents) {

	forgotPasswordRequest := ForgotPasswordRequest{}
	err := components.ValidateRequest(c, &forgotPasswordRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	exists := c.Components.Resources.Users.ExistsByEmail(forgotPasswordRequest.Email)

	if !exists {
		components.HttpResponse(c, http.StatusNoContent)
		return
	}

	randomToken := cacheutil.MustGenRandomToken()

	c.Components.Cache.Set(
		cacheutil.KeyWithPrefix(cacheutil.ForgotPasswordPrefix, randomToken),
		forgotPasswordRequest.Email,
		time.Minute*15,
	)

	updatePasswordURL := fmt.Sprintf(
		"%s:%s%s?t=%s",
		c.Components.Settings.String("frontend.host"),
		c.Components.Settings.String("frontend.port"),
		c.Components.Settings.String("frontend.updatePasswordEndpoint"),
		randomToken,
	)

	go c.Components.Mail.Send(
		[]string{forgotPasswordRequest.Email},
		mail.DefaultForgotPasswordSubject,
		fmt.Sprintf("Reset your password: %s", updatePasswordURL),
	)

	components.HttpResponse(c, http.StatusNoContent)
}

// UpdatePasswordHandler is the HTTP handler for updating the password
//
//	@Summary		Update password after email verification
//	@Description	Checks the token on the request and updates the password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			t		query	string					true	"User verification token"
//	@Param			request	body	UpdatePasswordRequest	true	"Update password info"
//	@Success		204		"No content"
//	@Failure		400		"Bad Request"
//
//	@Router			/auth/update-password [post]
func UpdatePasswordHandler(c *components.HTTPComponents) {
	randomToken := c.HttpRequest.URL.Query().Get("t")

	updatePasswordRequest := UpdatePasswordRequest{}
	err := components.ValidateRequest(c, &updatePasswordRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	email, found := c.Components.Cache.Get(
		cacheutil.KeyWithPrefix(cacheutil.ForgotPasswordPrefix, randomToken),
	)

	if !found {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUserResourceNotFound",
			}))
		return
	}

	c.Components.Cache.Delete(randomToken)

	user := &models.User{
		Email:    email.(string),
		Password: updatePasswordRequest.Password,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
		return
	}

	user.Password = string(hashedPassword)

	c.Components.Resources.Users.Update(user)

	components.HttpResponse(c, http.StatusNoContent)
}

// FirstAccessHandler is the HTTP handler checking if the user was pre-registered
//
//	@Summary		Checks if the user was pre-registered
//	@Description	Checks if the user was pre-registered and sends an e-mail to signup
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	FirstAccessRequest	true	"First access verification info"
//	@Success		204		"No content"
//	@Failure		400		"Bad Request"
//
//	@Router			/auth/first-access [post]
func FirstAccessHandler(c *components.HTTPComponents) {

	firstAccessRequest := FirstAccessRequest{}
	err := components.ValidateRequest(c, &firstAccessRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	found := c.Components.Resources.Users.ExistsDisabledByEmail(firstAccessRequest.Email)

	if found {
		randomToken := cacheutil.MustGenRandomToken()

		c.Components.Cache.Set(
			cacheutil.KeyWithPrefix(cacheutil.FirstAccessPrefix, randomToken),
			firstAccessRequest.Email,
			time.Minute*15,
		)

		updatePasswordURL := fmt.Sprintf(
			"%s:%s%s?t=%s",
			c.Components.Settings.String("frontend.host"),
			c.Components.Settings.String("frontend.port"),
			c.Components.Settings.String("frontend.firstAccessEndpoint"),
			randomToken,
		)

		go c.Components.Mail.Send(
			[]string{firstAccessRequest.Email},
			mail.DefaultFirstAccessSubject,
			fmt.Sprintf("Complete your signup: %s", updatePasswordURL),
		)
	}
	components.HttpResponse(c, http.StatusNoContent)
}

// FinishSignupHandler is the HTTP handler for filling up remaining user info
//
//	@Summary		Fills up remaining user info
//	@Description	Checks the token on the request and fills up remaining user info
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			t		query	string				true	"User verification token"
//	@Param			request	body	FinishSignupRequest	true	"Finish signup info"
//	@Success		204		"No content"
//	@Failure		400		"Bad Request"
//
//	@Router			/auth/finish-signup [post]
func FinishSignupHandler(c *components.HTTPComponents) {
	randomToken := c.HttpRequest.URL.Query().Get("t")

	finishSignupRequest := FinishSignupRequest{}
	err := components.ValidateRequest(c, &finishSignupRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	email, found := c.Components.Cache.Get(
		cacheutil.KeyWithPrefix(cacheutil.FirstAccessPrefix, randomToken),
	)
	if !found {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUserResourceNotFound",
			}))
		return
	}

	birthDate, _ := time.Parse(helpers.DefaultDateFormat(), finishSignupRequest.BirthDate)

	user := &models.User{
		Email:     email.(string),
		Name:      finishSignupRequest.Name,
		BirthDate: birthDate,
		CPF:       finishSignupRequest.CPF,
		Password:  finishSignupRequest.Password,
	}

	err = aws.HandleImageAndUploadToS3(
		finishSignupRequest.ProfilePicture,
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		components.HttpErrorLocalizedResponse(c, http.StatusBadRequest,
			c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ErrUnexpectedError",
			}))
		return
	}

	user.Password = string(hashedPassword)

	_, err = c.Components.Resources.Users.Update(user)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			components.HttpErrorLocalizedResponse(c, http.StatusNotFound,
				c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "ErrCPFAlreadyInUse",
				}))
			return
		} else {
			components.HttpErrorLocalizedResponse(c, http.StatusInternalServerError,
				c.Components.Localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "ErrUnexpectedError",
				}))
			return
		}
	}

	c.Components.Cache.Delete(randomToken)

	components.HttpResponse(c, http.StatusNoContent)
}
