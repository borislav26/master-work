package users

import (
	error2 "authentication-service/error"
	authentication_grpc "authentication-service/grpc/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserIsAuthenticated(authenticationService authentication_grpc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		request := &authentication_grpc.CheckUserAuthenticationRequest{
			UserToken: token,
		}
		res, err := authenticationService.CheckUserAuthentication(c, request)
		if err != nil {
			error2.HandleBadRequest(c, err, "Failed to check if user is authenticated")
			c.Abort()
			return
		}

		if !res.Result {
			error2.HandleBadRequestWithStatusCode(c, http.StatusUnauthorized, err, "Failed to check if user is authenticated")
			c.Abort()
			return
		}
		c.Next()
	}
}
