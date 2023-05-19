package helpers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type M map[string]any

func AssertHTTPResponse(t *testing.T, response *httptest.ResponseRecorder, expectedStatusCode int, expectedResponseBody map[string]any) {
	result := response.Result()
	defer result.Body.Close()

	responseBody := M{}
	err := json.NewDecoder(result.Body).Decode(&responseBody)

	assert.Nil(t, err, "invalid JSON in response body")
	assert.Equal(t, expectedStatusCode, result.StatusCode)
	assert.Equal(t, fmt.Sprint(responseBody), fmt.Sprint(expectedResponseBody))
}
