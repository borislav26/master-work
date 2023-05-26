package middlewares

import (
	error2 "cv-maker-service/error"
	"cv-maker-service/services/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckIfUserIsAuthenticated(authenticationService authentication.GRPCService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		request := &authentication.CheckUserAuthenticationRequest{
			UserToken: token,
		}

		res, err := authenticationService.CheckUserAuthentication(c, request)
		if err != nil {
			error2.HandleBadRequestWithStatusCode(c, http.StatusUnauthorized, err, "Failed to authenticate user")
			c.Abort()
			return
		}

		if !res.Result {
			error2.HandleBadRequestWithStatusCode(c, http.StatusUnauthorized, err, "Authorization failed")
			c.Abort()
			return
		}

		c.Next()
	}
}
