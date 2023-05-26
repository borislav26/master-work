package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"logger-service/environment"
)

func NewMongoDBManager() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(environment.MONGO_DATABASE_URL.Value()))
	if err != nil {
		return nil, err
	}

	return client, nil
}
