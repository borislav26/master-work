package generator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "pdf-generator-service/error"
	"pdf-generator-service/services/generator"
)

func getAllExamples(service generator.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		examples, err := service.GetAllExamples()
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to fetch pdf examples")
			return
		}
		c.JSON(http.StatusOK, examples)
	}
}
