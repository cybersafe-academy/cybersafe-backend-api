package helpers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertHTTPResponse(t *testing.T, response *httptest.ResponseRecorder, expectedStatusCode int, expectedResponseBody map[string]any) {
	result := response.Result()
	defer result.Body.Close()

	assert.Equal(t, expectedStatusCode, result.StatusCode)

	var responseBody map[string]any
	_ = json.NewDecoder(result.Body).Decode(&responseBody)

	assert.Equal(t, fmt.Sprint(responseBody), fmt.Sprint(expectedResponseBody))
}
