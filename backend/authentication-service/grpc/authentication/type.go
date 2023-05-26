package authentication_grpc

import (
	"authentication-service/database"
	"authentication-service/services/user"
	"context"
)

type (
	Service interface {
		CheckUserAuthentication(ctx context.Context, request *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error)
		GetUserEmailFromToken(ctx context.Context, request *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error)
		FindUserByEmail(ctx context.Context, request *FindUserByEmailRequest) (*FindUserByEmailResponse, error)
	}

	SimpleService struct {
		UnimplementedAuthenticationServiceServer
		DBManager   database.GormDbManager
		UserService user.Service
	}
)
