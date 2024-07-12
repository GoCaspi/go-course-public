package datasource_test

import (
	"example-project/datasource"
	"example-project/datasource/datasourcefakes"
	"example-project/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUpdateMany(t *testing.T) {
	fakeDb := &datasourcefakes.FakeMongoDBInterface{}
	fakeEmployees := []model.Employee{
		{ID: "1", FirstName: "First", LastName: "Last", Email: "Email"},
	}
	fakeDb.InsertManyReturns(&mongo.InsertManyResult{InsertedIDs: []interface{}{"1"}}, nil)
	dbClient := datasource.Client{
		Employee: fakeDb,
	}
	expectedData := make([]interface{}, len(fakeEmployees))
	expectedData[0] = fakeEmployees[0]
	err := dbClient.InsertEmployees(fakeEmployees)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, fakeDb.InsertManyCallCount())
	_, data, _ := fakeDb.InsertManyArgsForCall(0)
	assert.Equal(t, expectedData, data)
}
