package authentication

import (
	"authentication-service/database"
	error2 "authentication-service/error"
	"context"
)

func (s SimpleService) GenerateToken(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error) {
	return s.AuthService.GenerateToken(context, dbManager, userEmail)
}

func (s SimpleService) GenerateTokenFromPayload(payload map[string]any) (string, error2.Error) {
	return s.AuthService.GenerateTokenFromPayload(payload)
}

func (s SimpleService) FetchUserEmailFromToken(token string) (string, error2.Error) {
	return s.AuthService.FetchUserEmailFromToken(token)
}

func (s SimpleService) GenerateTokenForUserLogout(context context.Context, dbManager database.GormDbManager, userEmail string) (string, error2.Error) {
	return s.AuthService.GenerateTokenForUserLogout(context, dbManager, userEmail)
}
