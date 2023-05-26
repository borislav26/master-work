package makers

import (
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/maker"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func createCVFromProvidedData(
	dbManager database.GormDbManager,
	makerService maker.Service,
	authenticationService authentication.GRPCService,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *maker.CVData
		err := c.ShouldBindWith(&req, binding.JSON)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to bind request data")
			return
		}

		res, err := getUserEmailFromToken(c, authenticationService)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch user email for token")
			return
		}

		req.Email = res.Email

		err = makerService.CreateCV(c, dbManager, req)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to create cv from given data")
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
