package courses

import (
	"bytes"
	"context"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/internal/services"
	"cybersafe-backend-api/internal/services/reviews"
	"cybersafe-backend-api/pkg/helpers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/uuid"
)

func TestCreateReviewHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedResponseBody helpers.M
		payload              helpers.M
		queryParams          url.Values
		resourcesMock        services.Resources
	}{
		{
			name:               "success empty result",
			expectedStatusCode: 200,
			expectedResponseBody: helpers.M{
				"id":        uuid.Nil,
				"comment":   "lorem ipsulum",
				"rating":    1,
				"courseID":  "4f499382-dc42-43ae-af05-8a0411c32e03",
				"userID":    uuid.Nil,
				"createdAt": "0001-01-01T00:00:00Z",
				"deletedAt": nil,
				"updatedAt": "0001-01-01T00:00:00Z",
			},
			payload: helpers.M{
				"comment":  "lorem ipsulum",
				"rating":   1,
				"courseID": "4f499382-dc42-43ae-af05-8a0411c32e03",
				"userID":   "4f499382-dc42-43ae-af05-8a0411c32e03",
			},
			resourcesMock: services.Resources{
				Reviews: &reviews.ReviewsManagerMock{
					CreateMock: func(r *models.Review) error {
						return nil
					},
				},
			},
			queryParams: url.Values{
				"": []string{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			strPayload, _ := json.Marshal(testCase.payload)

			request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(strPayload))
			request.Header.Add("Content-Type", "application/json")
			request.URL.RawQuery = testCase.queryParams.Encode()

			ctx := context.WithValue(request.Context(), middlewares.UserKey, &models.User{})
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

			CreateCourseReview(httpComponentens)

			helpers.AssertHTTPResponse(t, response, testCase.expectedStatusCode, testCase.expectedResponseBody)
		})
	}
}
