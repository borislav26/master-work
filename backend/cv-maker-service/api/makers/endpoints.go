package makers

import (
	"cv-maker-service/api/middlewares"
	"cv-maker-service/database"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/maker"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	ctx *gin.Engine,
	dbManager database.GormDbManager,
	authenticationService authentication.GRPCService,
	makerService maker.Service,
) {
	group := ctx.Group("/api/maker")
	group.Use(middlewares.CheckIfUserIsAuthenticated(authenticationService))
	group.POST("/create-cv", createCVFromProvidedData(dbManager, makerService, authenticationService))
}
