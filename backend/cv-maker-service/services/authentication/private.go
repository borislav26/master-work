package authentication

import (
	"cv-maker-service/environement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s SimpleService) dialAuthenticationServiceConnectionAndReturnClient() (AuthenticationServiceClient, error) {
	conn, err := grpc.Dial(environement.GRPC_AUTHENTICATION_SERVICE_HOST.GetValue(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	authClient := NewAuthenticationServiceClient(conn)

	return authClient, nil
}
