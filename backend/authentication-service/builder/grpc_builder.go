package builder

import (
	"authentication-service/database"
	"authentication-service/environement"
	authentication_grpc "authentication-service/grpc/authentication"
	"google.golang.org/grpc"
	"net"
)

type GRPCServices struct {
	AuthenticationService *authentication_grpc.SimpleService
}

func BuildGRPCServices(dbManager database.DBManager, services *Services) *GRPCServices {
	authenticationService := &authentication_grpc.SimpleService{
		DBManager:   dbManager.GORMDBManager,
		UserService: services.UserService,
	}
	return &GRPCServices{
		AuthenticationService: authenticationService,
	}
}

func ServeGRPCServer(grpcServices *GRPCServices) error {
	listener, err := net.Listen("tcp", environement.GRPC_SERVER_HOST.GetValue())
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	authentication_grpc.RegisterAuthenticationServiceServer(grpcServer, grpcServices.AuthenticationService)

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
