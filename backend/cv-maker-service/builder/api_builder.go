package builder

import (
	"cv-maker-service/api/cv_data_byte"
	"cv-maker-service/api/languages"
	"cv-maker-service/api/makers"
	"cv-maker-service/api/programming_languages"
	"cv-maker-service/database"
	"github.com/gin-gonic/gin"
)

func BuildApiLayer(dbManager database.GormDbManager, services *Services, engine *gin.Engine, emitter *Emitter) {
	//config := cors.DefaultConfig()
	//config.AllowHeaders = append(config.AllowHeaders, "*")
	//config.AllowAllOrigins = true
	//engine.Use(cors.New(config))

	programming_languages.AddEndpoints(engine, dbManager, services.ProgrammingLanguageService, services.GRPCAuthenticationService)
	languages.AddEndpoints(engine, dbManager, services.LanguageService, services.GRPCAuthenticationService)
	makers.AddEndpoints(engine, dbManager, services.GRPCAuthenticationService, services.MakerService)
	cv_data_byte.AddEndpoints(engine, dbManager, services.CVDataByteService, services.GRPCAuthenticationService)
}
