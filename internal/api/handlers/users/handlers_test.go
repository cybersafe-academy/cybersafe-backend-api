package users

import (
	"bytes"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/internal/services"
	"cybersafe-backend-api/internal/services/users"
	"cybersafe-backend-api/pkg/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsersHandler(t *testing.T) {
	cases := []struct {
		testName             string
		expectedStatusCode   int
		expectedResponseBody map[string]any
	}{
		{
			testName:           "success empty result",
			expectedStatusCode: 200,
			expectedResponseBody: map[string]any{
				"data":       nil,
				"total":      0,
				"limit":      10,
				"current":    1,
				"totalPages": 0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {

			payload := bytes.NewBuffer(nil)

			request := httptest.NewRequest(http.MethodGet, "/users", payload)
			response := httptest.NewRecorder()

			request.Header.Add("Content-Type", "application/json")

			c := &components.Components{
				Resources: services.Resources{
					Users: &users.UsersManagerMock{
						ListWithPaginationMock: func(limit, offset int) ([]models.User, int64) {
							return []models.User{}, 0
						},
					},
				},
			}

			httpComponentens := &components.HTTPComponents{
				Components:   c,
				HttpRequest:  request,
				HttpResponse: response,
			}

			ListUsersHandler(httpComponentens)

			helpers.AssertHTTPResponse(t, response, tc.expectedStatusCode, tc.expectedResponseBody)
		})
	}
}
