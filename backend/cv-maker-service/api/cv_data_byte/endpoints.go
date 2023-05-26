package cv_data_byte

import (
	"cv-maker-service/api/middlewares"
	"cv-maker-service/database"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/cv_data_byte"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	ctx *gin.Engine,
	dbManager database.GormDbManager,
	cvDataByteService cv_data_byte.Service,
	authenticationService authentication.GRPCService,
) {
	group := ctx.Group("/api/cv-data-byte")
	group.Use(middlewares.CheckIfUserIsAuthenticated(authenticationService))
	group.GET("/all-for-user", getAllForUser(dbManager, cvDataByteService, authenticationService))
}
