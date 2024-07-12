package handler_test

import (
	"encoding/json"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(responseRecoder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "1"})

	fakeEmployee := model.Employee{
		ID:        "1",
		FirstName: "Joe",
	}
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(&fakeEmployee, nil)

	expectedBody, _ := json.Marshal(fakeEmployee)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContext)

	assert.Equal(t, http.StatusOK, responseRecoder.Code)
	assert.Equal(t, expectedBody, responseRecoder.Body.Bytes())

}
