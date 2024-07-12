package service_test

import (
	"example-project/model"
	"example-project/service"
	"example-project/service/servicefakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}
	fakeDB.GetEmployeeByIDReturns(&data, nil)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, err := serviceInstance.GetEmployeeById("1")
	assert.Equal(t, nil, err)
	assert.Equal(t, &data, actual)
}
