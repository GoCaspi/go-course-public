package datasource

import (
	"context"
	"errors"
	"example-project/errorHandler"
	"example-project/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const idField = "id"

func (c *Client) InsertEmployees(employees []model.Employee) error {
	entities := make([]interface{}, len(employees))
	for index, entity := range employees {
		entities[index] = entity
	}
	_, err := c.Employee.InsertMany(context.Background(), entities)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New(errorHandler.InternalErrorEmployeeIdNotUnique)
		}
		return errors.New(errorHandler.InternalErrorDatabase)
	}
	return nil
}

func (c *Client) GetEmployeeByID(id string) (*model.Employee, error) {
	filter := bson.M{idField: id}
	cursor := c.Employee.FindOne(context.TODO(), filter)
	if cursor.Err() != nil {
		if errors.Is(cursor.Err(), mongo.ErrNoDocuments) {
			return nil, errors.New(errorHandler.InternalErrorNoEmployeeFound)
		}
		return nil, errors.New(errorHandler.InternalErrorDatabase)
	}
	var employee model.Employee
	err := cursor.Decode(&employee)
	if err != nil {
		return nil, errors.New(errorHandler.InternalErrorMarshaling)
	}
	return &employee, nil
}
