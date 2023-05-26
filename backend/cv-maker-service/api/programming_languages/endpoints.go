package programming_languages

import (
	"cv-maker-service/api/middlewares"
	"cv-maker-service/database"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/programming_language"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	ctx *gin.Engine,
	dbManager database.GormDbManager,
	programmingLanguageService programming_language.Service,
	authenticationService authentication.GRPCService,
) {
	group := ctx.Group("/api/programming_languages")
	group.Use(middlewares.CheckIfUserIsAuthenticated(authenticationService))
	group.GET("/all", getAllProgrammingLanguages(dbManager, programmingLanguageService))
}
