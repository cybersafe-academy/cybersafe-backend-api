package middlewares

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/internal/services"
	"cybersafe-backend-api/internal/services/users"
	"cybersafe-backend-api/pkg/cacheutil"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/jwtutil"
	"cybersafe-backend-api/pkg/settings"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthorization(t *testing.T) {
	testCases := []struct {
		name                 string
		allowedRoles         []string
		expectedStatusCode   int
		header               http.Header
		expectedResponseBody helpers.M
		components           components.Components
	}{
		{
			name:               "header authorization is empty",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        http.StatusUnauthorized,
					"description": errutil.ErrCredentialsMissing.Error(),
				},
			},
		},
		{
			name:               "token malformed",
			expectedStatusCode: http.StatusUnauthorized,
			header: http.Header{
				"Authorization": []string{"foo"},
			},
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        http.StatusUnauthorized,
					"description": errutil.ErrInvalidJWT.Error(),
				},
			},
			components: components.Components{
				Settings: &settings.SettingsMock{
					Source: helpers.M{
						"jwt.secretKey": "secretKey",
					},
				},
			},
		},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			router := chi.NewRouter()
			router.Use(Authorizer(&testCase.components, testCase.allowedRoles...))
			router.Get("/", testHandler)

			request := httptest.NewRequest("GET", "/", nil)
			request.Header = testCase.header

			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}

func TestAuthorizationToken(t *testing.T) {
	testCases := []struct {
		name                 string
		allowedRoles         []string
		jwtExpiresAt         time.Duration
		jwtID                uuid.UUID
		expectedStatusCode   int
		expectedResponseBody helpers.M
		user                 models.User
		components           components.Components
	}{
		{
			name:               "user referenced in token is not found",
			expectedStatusCode: http.StatusNotFound,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        http.StatusNotFound,
					"description": errutil.ErrUserResourceNotFound.Error(),
				},
			},
			jwtExpiresAt: 1 * time.Hour,
			jwtID:        uuid.New(),
			user: models.User{
				Shared: models.Shared{
					ID: uuid.New(),
				},
				Role: models.DefaultUserRole,
			},
			components: components.Components{
				Settings: &settings.SettingsMock{
					Source: helpers.M{
						"jwt.secretKey":      "secretKey",
						"jwt.subject":        "userCredentials",
						"application.issuer": "cybersafe",
					},
				},
				Cache: &cacheutil.CacheMock{
					Source: helpers.M{},
				},
				Resources: services.Resources{
					Users: &users.UsersManagerMock{
						GetByIDMock: func(u uuid.UUID) (models.User, error) {
							return models.User{}, gorm.ErrRecordNotFound
						},
					},
				},
			},
		},
		{
			name:               "unexpected error while trying to retrieve user",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        http.StatusInternalServerError,
					"description": errutil.ErrUnexpectedError.Error(),
				},
			},
			jwtExpiresAt: 1 * time.Hour,
			jwtID:        uuid.New(),
			user: models.User{
				Shared: models.Shared{
					ID: uuid.New(),
				},
				Role: models.DefaultUserRole,
			},
			components: components.Components{
				Settings: &settings.SettingsMock{
					Source: helpers.M{
						"jwt.secretKey":      "secretKey",
						"jwt.subject":        "userCredentials",
						"application.issuer": "cybersafe",
					},
				},
				Cache: &cacheutil.CacheMock{
					Source: helpers.M{},
				},
				Resources: services.Resources{
					Users: &users.UsersManagerMock{
						GetByIDMock: func(u uuid.UUID) (models.User, error) {
							return models.User{}, gorm.ErrInvalidData
						},
					},
				},
			},
		},
		{
			name:               "given user role is not allowed",
			allowedRoles:       []string{models.AdminUserRole},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        http.StatusInternalServerError,
					"description": errutil.ErrUnexpectedError.Error(),
				},
			},
			jwtExpiresAt: 1 * time.Hour,
			jwtID:        uuid.New(),
			user: models.User{
				Shared: models.Shared{
					ID: uuid.New(),
				},
				Role: models.DefaultUserRole,
			},
			components: components.Components{
				Settings: &settings.SettingsMock{
					Source: helpers.M{
						"jwt.secretKey":      "secretKey",
						"jwt.subject":        "userCredentials",
						"application.issuer": "cybersafe",
					},
				},
				Cache: &cacheutil.CacheMock{
					Source: helpers.M{},
				},
				Resources: services.Resources{
					Users: &users.UsersManagerMock{
						GetByIDMock: func(u uuid.UUID) (models.User, error) {
							return models.User{}, nil
						},
					},
				},
			},
		},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			router := chi.NewRouter()
			router.Use(Authorizer(&testCase.components, testCase.allowedRoles...))
			router.Get("/", testHandler)

			request := httptest.NewRequest("GET", "/", nil)

			jwtClaims := jwtutil.CustomClaims{
				UserID: testCase.user.ID.String(),
				Role:   testCase.user.Role,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    testCase.components.Settings.String("application.issuer"),
					Subject:   testCase.components.Settings.String("jwt.subject"),
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(testCase.jwtExpiresAt)),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
					NotBefore: jwt.NewNumericDate(time.Now()),
					ID:        testCase.jwtID.String(),
				},
			}

			tokenString, err := jwtutil.Generate(jwtClaims, testCase.components.Settings.String("jwt.secretKey"))

			assert.NoError(t, err, "failed to generate token")

			request.Header.Add("Authorization", tokenString)

			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}
