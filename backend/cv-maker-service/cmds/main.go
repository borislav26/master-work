package main

import (
	"cv-maker-service/builder"
	"cv-maker-service/database"
	"cv-maker-service/environement"
	"github.com/gin-gonic/gin"
)

func main() {
	dbManager := database.NewDBManager()
	defer dbManager.CloseConnection()

	//channel := builder.GetChannelAndConnectToRabbitMQ()

	repositories := builder.BuildRepositories()

	services := builder.BuildServices(repositories)

	app := gin.Default()

	builder.BuildApiLayer(dbManager.GORMDBManager, services, app, nil)

	app.Run(environement.HTTP_SERVER_PORT.GetValue())
}
