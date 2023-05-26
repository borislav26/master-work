package authentication

import (
	"authentication-service/database"
	error2 "authentication-service/error"
	"authentication-service/services/user"
	"context"
)

const (
	AlgorithmKey  string = "alg"
	TypeKey       string = "type"
	ExpirationKey string = "exp"
	UserEmailKey  string = "userEmail"
)

type (
	Service interface {
		GenerateToken(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error)
		GenerateTokenFromPayload(payload map[string]any) (string, error2.Error)
		FetchUserEmailFromToken(token string) (string, error2.Error)
		GenerateTokenForUserLogout(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error)
	}

	SimpleService struct {
		AuthService
	}
)

var HeaderJSON map[string]any

func init() {
	HeaderJSON = map[string]any{
		"alg": "SHA256",
		"typ": "JWT",
	}
}

func NewService(userService user.Service) Service {
	return &SimpleService{
		AuthService: NewAuthService(userService),
	}
}
