package middlewares

import (
	"github.com/gin-gonic/gin"
	"pdf-generator-service/services/authentication"
)

func CheckIfUserIsAuthenticated(authenticationService authentication.GRPCService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
