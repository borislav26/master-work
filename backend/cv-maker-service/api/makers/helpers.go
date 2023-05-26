package makers

import (
	"cv-maker-service/services/authentication"
	"github.com/gin-gonic/gin"
)

func getUserEmailFromToken(c *gin.Context, service authentication.GRPCService) (*authentication.GetUserEmailFromTokenResponse, error) {
	token := c.Request.Header.Get("Authorization")
	req := &authentication.CheckUserAuthenticationRequest{
		UserToken: token,
	}
	res, err := service.GetUserEmailFromToken(c, req)
	if err != nil {
		return res, err
	}

	return res, nil
}
