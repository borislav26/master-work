package main

import (
	"log"
	"logger-service/builder"
	"logger-service/repo/database"
)

func main() {
	mongoClient, err := database.NewMongoDBManager()
	if err != nil {
		log.Fatalln("Failed to connect to mongodb service")
	}

	repositories := builder.BuildRepositories(mongoClient)
	services := builder.BuildServices(repositories)

	err = builder.ServeGRPCServer(services)
	if err != nil {
		log.Fatalln(err)
	}
}
