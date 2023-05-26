package main

import (
	"authentication-service/builder"
	"authentication-service/database"
	"authentication-service/environement"
	"github.com/gin-gonic/gin"
)

func main() {
	dbManager := database.NewDBManager()
	defer dbManager.CloseConnection()

	repositories := builder.BuildRepositories()

	services := builder.BuildServices(repositories)

	grpcServices := builder.BuildGRPCServices(dbManager, services)

	app := gin.Default()

	builder.BuildApiLayer(dbManager.GORMDBManager, services, app, grpcServices)

	go builder.ServeGRPCServer(grpcServices)

	app.Run(environement.HTTP_SERVER_PORT.GetValue())
}
