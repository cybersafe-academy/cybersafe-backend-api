package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/jwtutil"
	"errors"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

	user := models.User{}

	result := db.MustGetDbConn().Where("CPF = ?", loginRequest.CPF).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUserResourceNotFound)
			return
		} else {
			components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	expirationTime := 1 * time.Hour
	JWTID := uuid.NewString()

	claims := jwtutil.CustomClaims{

		UserID: user.ID.String(),
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
		time.Until(jwtClaims.RegisteredClaims.ExpiresAt.Time),
		jwtClaims.RegisteredClaims.ID,
		token.Raw,
	)

	components.HttpResponse(c, http.StatusNoContent)
}
