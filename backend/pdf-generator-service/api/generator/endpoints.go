package generator

import (
	"github.com/gin-gonic/gin"
	"pdf-generator-service/api/middlewares"
	"pdf-generator-service/services/authentication"
	"pdf-generator-service/services/generator"
)

func AddEndpoints(engine *gin.Engine, service generator.Service, authenticationService authentication.GRPCService) {
	group := engine.Group("/api/web/generator")
	group.Use(middlewares.CheckIfUserIsAuthenticated(authenticationService))
	group.GET("/examples", getAllExamples(service))
}
