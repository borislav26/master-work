package languages

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/services/language"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllLanguages(dbManager database.GormDbManager, languageService language.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		languages, err := languageService.All(dbManager)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch all languages")
			return
		}

		c.JSON(http.StatusOK, languages)
	}
}
