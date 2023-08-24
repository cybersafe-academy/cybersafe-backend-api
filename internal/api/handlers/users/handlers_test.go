package users

import (
	"bytes"
	"context"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/internal/services"
	"cybersafe-backend-api/internal/services/users"
	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/validation"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/uuid"
)

func TestListUsersHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedResponseBody helpers.M
		queryParams          url.Values
		resourcesMock        services.Resources
	}{
		{
			name:               "success empty result",
			expectedStatusCode: 200,
			expectedResponseBody: helpers.M{
				"data":       nil,
				"total":      0,
				"limit":      10,
				"current":    1,
				"totalPages": 0,
			},
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					ListWithPaginationMock: func(limit, offset int) ([]models.User, int64) {
						return []models.User{}, 0
					},
				},
			},
		},
		{
			name:               "success non-empty result",
			expectedStatusCode: 200,
			expectedResponseBody: helpers.M{
				"data": []helpers.M{
					{
						"birthDate": "0001-01-01",
						"cpf":       "",
						"createdAt": "0001-01-01T00:00:00Z",
						"deletedAt": nil,
						"email":     "",
						"id":        uuid.Nil,
						"name":      "Test user",
						"role":      "",
						"updatedAt": "0001-01-01T00:00:00Z",
					},
				},
				"total":      1,
				"limit":      10,
				"current":    1,
				"totalPages": 1,
			},
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					ListWithPaginationMock: func(limit, offset int) ([]models.User, int64) {
						users := []models.User{
							{
								Name: "Test user",
							},
						}
						return users, int64(len(users))
					},
				},
			},
		},
		{
			name:               "invalid query params page",
			expectedStatusCode: 400,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        400,
					"description": "invalid page param",
				},
			},
			queryParams: url.Values{
				"page": []string{"invalid"},
			},
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					ListWithPaginationMock: func(limit, offset int) ([]models.User, int64) {
						return []models.User{}, 0
					},
				},
			},
		},
		{
			name:               "invalid query params limit",
			expectedStatusCode: 400,
			expectedResponseBody: helpers.M{
				"error": helpers.M{
					"code":        400,
					"description": "invalid limit param",
				},
			},
			queryParams: url.Values{
				"limit": []string{"invalid"},
			},
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					ListWithPaginationMock: func(limit, offset int) ([]models.User, int64) {
						return []models.User{}, 0
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			payload := bytes.NewBuffer(nil)

			request := httptest.NewRequest(http.MethodGet, "/users", payload)
			request.Header.Add("Content-Type", "application/json")
			request.URL.RawQuery = testCase.queryParams.Encode()

			response := httptest.NewRecorder()
			c := &components.Components{
				Resources: testCase.resourcesMock,
			}

			httpComponentens := &components.HTTPComponents{
				Components:   c,
				HttpRequest:  request,
				HttpResponse: response,
			}

			ListUsersHandler(httpComponentens)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}

func TestCreateUserHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedResponseBody helpers.M
		queryParams          url.Values
		payload              helpers.M
		resourcesMock        services.Resources
	}{
		{
			name:               "success empty result",
			expectedStatusCode: 200,
			expectedResponseBody: helpers.M{
				"birthDate": "0001-01-01",
				"cpf":       "12427904970",
				"createdAt": "0001-01-01T00:00:00Z",
				"deletedAt": nil,
				"email":     "teste@email.com",
				"id":        uuid.Nil,
				"name":      "Test User",
				"role":      models.DefaultUserRole,
				"updatedAt": "0001-01-01T00:00:00Z",
			},
			payload: helpers.M{
				"birthDate": "0001-01-01",
				"cpf":       "12427904970",
				"email":     "teste@email.com",
				"name":      "Test User",
				"password":  "12345678",
				"role":      models.DefaultUserRole,
			},
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					CreateMock: func(u *models.User) error {
						return nil
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			payload := helpers.MustMapToBytesBuffer(testCase.payload)

			request := httptest.NewRequest(http.MethodPost, "/users", payload)
			request.Header.Add("Content-Type", "application/json")
			request.URL.RawQuery = testCase.queryParams.Encode()

			ctx := context.WithValue(
				request.Context(),
				middlewares.UserKey,
				&models.User{
					Role: models.MasterUserRole,
				},
			)
			request = request.WithContext(ctx)

			response := httptest.NewRecorder()
			c := &components.Components{
				Resources: testCase.resourcesMock,
			}

			httpComponentens := &components.HTTPComponents{
				Components:   c,
				HttpRequest:  request,
				HttpResponse: response,
			}

			validation.Config()

			CreateUserHandler(httpComponentens)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}

func TestPersonalityTestHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedResponseBody helpers.M
		queryParams          url.Values
		resourcesMock        services.Resources
		mbtiType             mbti
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusNoContent,
			mbtiType:           INTJ,
			resourcesMock: services.Resources{
				Users: &users.UsersManagerMock{
					UpdateMock: func(user *models.User) (int, error) {
						return 1, nil
					},
				},
			},
		},
		{
			name:               "failed because invalid mbti type",
			expectedStatusCode: 400,
			expectedResponseBody: map[string]any{
				"error": map[string]any{
					"code":        400,
					"description": "invalid MBTI type",
				},
			},
			queryParams:   map[string][]string{},
			resourcesMock: services.Resources{},
			mbtiType:      "INVALID",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			requestMap := map[string]any{
				"mbtiType": testCase.mbtiType,
			}

			payload := helpers.MustMapToBytesBuffer(requestMap)

			request := httptest.NewRequest(http.MethodPost, "/personality-test", payload)
			request.Header.Add("Content-Type", "application/json")
			request.URL.RawQuery = testCase.queryParams.Encode()
			ctx := context.WithValue(request.Context(), middlewares.UserKey, &models.User{})
			request = request.WithContext(ctx)

			response := httptest.NewRecorder()
			c := &components.Components{
				Resources: testCase.resourcesMock,
			}

			httpComponents := &components.HTTPComponents{
				Components:   c,
				HttpRequest:  request,
				HttpResponse: response,
			}

			PersonalityTestHandler(httpComponents)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}
