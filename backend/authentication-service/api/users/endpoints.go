package users

import (
	"authentication-service/database"
	authentication_grpc "authentication-service/grpc/authentication"
	"authentication-service/services/authentication"
	"authentication-service/services/user"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	ctx *gin.Engine,
	dbManager database.GormDbManager,
	userService user.Service,
	authenticationService authentication.Service,
	grpcAuthenticationService authentication_grpc.Service,
) {
	group := ctx.Group("/api/users")
	group.POST("/sign-up", signUpNewUser(dbManager, userService, authenticationService))
	group.POST("/login", login(dbManager, userService, authenticationService))
	group.POST("/log-out", logout(dbManager, authenticationService))

	authenticatedGroup := ctx.Group("/api/web/users")
	authenticatedGroup.Use(CheckUserIsAuthenticated(grpcAuthenticationService))
	authenticatedGroup.GET("/me", getUserMe(dbManager, userService, authenticationService))
}
