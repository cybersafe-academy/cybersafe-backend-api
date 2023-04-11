package authentication

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/models"
	"errors"
	"time"

	"net/http"

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
			components.HttpErrorResponse(c, http.StatusBadRequest, errutil.ErrUnexpectedError)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		components.HttpErrorResponse(c, http.StatusInternalServerError, errutil.ErrUnexpectedError)
		return
	}

	expirationTime := 24 * time.Hour

	tokenString, err := jwtutil.Generate(
		user.ID.String(),
		c.Components.Settings.String("application.issuer"),
		c.Components.Settings.String("jwt.subject"),
		c.Components.Settings.String("jwt.secretKey"),
		expirationTime,
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
