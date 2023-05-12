package users

import (
	"bytes"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/settings"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsersHandler(t *testing.T) {
	cases := []struct {
		testName                      string
		expectedStatusCode            int
		expectedResponseErrorMessage  string
		expectedResponseUserTextAlert string
	}{}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {

			payload := bytes.NewBuffer(nil)

			request := httptest.NewRequest(http.MethodPut, "/activate", payload)
			response := httptest.NewRecorder()

			request.Header.Add("Content-Type", "application/json")

			// rctx := chi.NewRouteContext()

			c := &components.Components{
				Settings: &settings.SettingsMock{
					Source: map[string]any{
						"teste": 1,
					},
				},
			}

			httpComponentens := &components.HTTPComponents{
				Components:   c,
				HttpRequest:  request,
				HttpResponse: response,
			}

			ListUsersHandler(httpComponentens)

			helpers.AssertHTTPResponse(t, response, tc.expectedStatusCode)
		})
	}
}
