package builder

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pdf-generator-service/api/generator"
)

func BuildAPILayer(engine *gin.Engine, services *Services) {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "*")
	config.AllowAllOrigins = true
	//config.AllowOrigins = append(config.AllowOrigins, "http://frontend-service:3000", "http://localhost:3000")
	engine.Use(cors.New(config))
	generator.AddEndpoints(engine, services.GeneratorService, services.AuthenticationGRPCService)
}
