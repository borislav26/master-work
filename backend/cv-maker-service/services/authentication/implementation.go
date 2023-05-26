package authentication

import (
	"context"
	error2 "cv-maker-service/error"
)

func (s SimpleService) CheckUserAuthentication(context context.Context, req *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error) {
	authClient, err := s.dialAuthenticationServiceConnectionAndReturnClient()
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}
	res, err := authClient.CheckUserAuthentication(context, req)
	return res, err
}

func (s SimpleService) GetUserEmailFromToken(ctx context.Context, request *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error) {
	authClient, err := s.dialAuthenticationServiceConnectionAndReturnClient()
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}
	res, err := authClient.GetUserEmailFromToken(ctx, request)
	return res, err
}

func (s SimpleService) FindUserByEmail(ctx context.Context, request *FindUserByEmailRequest) (*FindUserByEmailResponse, error) {
	authClient, err := s.dialAuthenticationServiceConnectionAndReturnClient()
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}
	res, err := authClient.FindUserByEmail(ctx, request)
	return res, err
}
