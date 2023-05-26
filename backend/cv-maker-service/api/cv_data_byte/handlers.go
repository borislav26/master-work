package cv_data_byte

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/cv_data_byte"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllForUser(dbManager database.GormDbManager, cvDataByteService cv_data_byte.Service, authService authentication.GRPCService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := getUserEmailFromToken(c, authService)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch user email for token")
			return
		}

		cvs, err := cvDataByteService.GetAllGeneratedCVsForUser(c, dbManager, email.Email)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch all languages")
			return
		}

		c.JSON(http.StatusOK, cvs)
	}
}
