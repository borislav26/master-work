package builder

import (
	"authentication-service/api/users"
	"authentication-service/database"
	"github.com/gin-gonic/gin"
)

func BuildApiLayer(dbManager database.GormDbManager, services *Services, engine *gin.Engine, grpcServices *GRPCServices) {
	//engine.Use(users.SetHeader("Access-Control-Allow-Origin", "*"))
	//engine.Use(users.SetHeader("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE"))
	//engine.Use(users.SetHeader("Access-Control-Allow-Credentials", "true"))
	//engine.Use(users.SetHeader("Content-Type", "application/json"))
	//engine.Use(users.SetHeader("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding"))
	//config := cors.DefaultConfig()
	//config.AllowHeaders = append(config.AllowHeaders, "*")
	//config.AllowOrigins = append(config.AllowOrigins, "http://frontend-service:3000", "http://localhost:3000", "http://frontend-service")
	//engine.Use(cors.New(config))

	//engine.Use(users.SetHeader("Access-Control-Allow-Origin", "*"))
	users.AddEndpoints(engine, dbManager, services.UserService, services.AuthenticationService, grpcServices.AuthenticationService)
}
