package programming_languages

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/services/programming_language"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllProgrammingLanguages(dbManager database.GormDbManager, programmingLanguageService programming_language.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		programmingLanguages, err := programmingLanguageService.All(dbManager)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch all programming languages")
			return
		}

		c.JSON(http.StatusOK, programmingLanguages)
	}
}
