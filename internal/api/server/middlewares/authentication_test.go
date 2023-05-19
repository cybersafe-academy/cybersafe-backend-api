package middlewares

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/settings"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
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
		{
			name:               "token malformed",
			expectedStatusCode: http.StatusUnauthorized,
			header: http.Header{
				"Authorization": []string{},
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
						"jwt.secretKey": "foo",
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
