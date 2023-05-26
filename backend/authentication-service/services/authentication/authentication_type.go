package authentication

import (
	"authentication-service/database"
	error2 "authentication-service/error"
	"authentication-service/services/user"
	"context"
)

type (
	AuthService interface {
		GenerateToken(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error)
		GenerateTokenFromPayload(payload map[string]any) (string, error2.Error)
		FetchUserEmailFromToken(token string) (string, error2.Error)
		GenerateTokenForUserLogout(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error)
	}

	AuthSimpleService struct {
		UserService user.Service
	}
)

func NewAuthService(userService user.Service) AuthService {
	return &AuthSimpleService{
		UserService: userService,
	}
}
