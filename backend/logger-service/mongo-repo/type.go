package mongo_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		SaveLog(context context.Context, record any) error
	}

	SimpleRepository struct {
		collection *mongo.Collection
	}
)

func NewSimpleRepository(client *mongo.Client) *SimpleRepository {
	collection := client.Database("master_work").Collection("logs")
	return &SimpleRepository{
		collection: collection,
	}
}
