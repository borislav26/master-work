package authentication

import "context"

type (
	GRPCService interface {
		CheckUserAuthentication(context.Context, *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error)
		GetUserEmailFromToken(ctx context.Context, request *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error)
		FindUserByEmail(ctx context.Context, request *FindUserByEmailRequest) (*FindUserByEmailResponse, error)
	}

	SimpleService struct {
	}
)
