package builder

import (
	"go.mongodb.org/mongo-driver/mongo"
	"logger-service/mongo-repo"
)

type Repositories struct {
	MongoRepository mongo_repo.Repository
}

func BuildRepositories(client *mongo.Client) *Repositories {
	return &Repositories{
		MongoRepository: mongo_repo.NewSimpleRepository(client),
	}
}
