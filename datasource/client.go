package datasource

import (
	"context"
	"example-project/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MongoDBInterface
type MongoDBInterface interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
}

type Client struct {
	Employee MongoDBInterface
}

func NewDbClient(d model.DbConfig) (*Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.URL))
	if err != nil {
		return nil, err
	}
	db := client.Database(d.Database)
	return &Client{
		Employee: db.Collection(employeeCollection),
	}, nil
}
