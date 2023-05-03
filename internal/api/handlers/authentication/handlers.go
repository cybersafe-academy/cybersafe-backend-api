package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/mail"
	"fmt"
	"time"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	user, err := c.Components.Storer.GetUserByCPF(loginRequest.CPF)

	if err != nil {
		components.HttpErrorResponse(c, http.StatusUnauthorized, errutil.ErrUserResourceNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		components.HttpErrorResponse(c, http.StatusUnauthorized, errutil.ErrLoginOrPasswordIncorrect)
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
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
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
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
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
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
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
	token, _ := jwtutil.Parse(authorizationHeader, jwtClaims)

	jwtutil.AddToBlackList(
		c.Components.Cache,
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

	if err := c.Components.DB.Where("email = ?", forgotPasswordRequest.Email).First(&models.User{}).Error; err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUserResourceNotFound)
		return
	}

	randomToken := helpers.MustGenerateURLEncodedRandomToken()

	(*c.Components.Cache).Set(
		randomToken, forgotPasswordRequest.Email, time.Minute*15,
	)

	updatePasswordURL := fmt.Sprintf(
		"%s:%s%s%s?t=%s",
		c.Components.Settings.String("docs.host"),
		c.Components.Settings.String("docs.port"),
		c.Components.Settings.String("docs.basePath"),
		c.Components.Settings.String("mail.updatePasswordEndpoint"),
		randomToken,
	)

	(*c.Components.Mail).Send(
		[]string{forgotPasswordRequest.Email},
		mail.DefaultForgotPasswordSubject,
		fmt.Sprintf("Reset your password: %s", updatePasswordURL),
	)

	components.HttpResponse(c, http.StatusNoContent)
}

// ForgotPasswordHandler is the HTTP handler for requesting a new password
//
//	@Summary		Update password after email verification
//	@Description	Checks the token on the request and updates the password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			t	query	string	true	"User verification token"
//
//	@Success		204	"No content"
//	@Failure		400	"Bad Request"
//
//	@Router			/auth/update-password [post]
func UpdatePasswordHandler(c *components.HTTPComponents) {
	randomToken := chi.URLParam(c.HttpRequest, "t")

	updatePasswordRequest := UpdatePasswordRequest{}
	err := components.ValidateRequest(c, &updatePasswordRequest)
	if err != nil {
		components.HttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	email, found := (*c.Components.Cache).Get(randomToken)

	if !found {
		components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUserResourceNotFound)
		return
	}

	user := &models.User{
		Email:    email.(string),
		Password: updatePasswordRequest.Password,
	}

	c.Components.DB.Model(&models.User{}).Where("email = ?", email).Updates(user)

	components.HttpResponse(c, http.StatusNoContent)
}
