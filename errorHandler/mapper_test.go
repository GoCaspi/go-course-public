package errorHandler_test

import (
	"example-project/errorHandler"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestHandleError(t *testing.T) {
	tests := []struct {
		name           string
		intErr         string
		expectedErr    string
		expectedStatus int
	}{
		{"case 1", errorHandler.InternalErrorDatabase, errorHandler.ExternalErrorUnknown, 500},
		{"case 2", errorHandler.InternalErrorMarshaling, errorHandler.ExternalErrorUnknown, 500},
		{"case 3", errorHandler.InternalErrorNoEmployeeFound, errorHandler.ExternalErrorEmployeeNotFound, 404},
		{"case 4", errorHandler.InternalErrorInvalidPayload, errorHandler.ExternalErrorInvalidPayload, 400},
		{"case 5", errorHandler.InternalErrorEmployeeIdNotGiven, errorHandler.ExternalErrorIdMissingInURL, 400},
		{"case 6", errorHandler.InternalErrorEmployeeIdNotUnique, errorHandler.ExternalErrorEmployeeIdNotUnique, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecoder := httptest.NewRecorder()
			fakeContext, _ := gin.CreateTestContext(responseRecoder)
			errorHandler.HandleError(fakeContext, tt.intErr)
			expected := `{"errorMessage":"` + tt.expectedErr + `"}`
			assert.Equal(t, expected, responseRecoder.Body.String())
			assert.Equal(t, tt.expectedStatus, responseRecoder.Code)
		})
	}
}
