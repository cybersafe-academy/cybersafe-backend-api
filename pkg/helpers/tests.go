package helpers

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertHTTPResponse(t *testing.T, res *httptest.ResponseRecorder, expectedStatusCode int) {
	result := res.Result()
	defer result.Body.Close()

	assert.Equal(t, expectedStatusCode, result.StatusCode)
}
