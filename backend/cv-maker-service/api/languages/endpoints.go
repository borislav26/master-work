package languages

import (
	"cv-maker-service/api/middlewares"
	"cv-maker-service/database"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/language"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	ctx *gin.Engine,
	dbManager database.GormDbManager,
	programmingLanguageService language.Service,
	authenticationService authentication.GRPCService,
) {
	group := ctx.Group("/api/languages")
	group.Use(middlewares.CheckIfUserIsAuthenticated(authenticationService))
	group.GET("/all", getAllLanguages(dbManager, programmingLanguageService))
}
