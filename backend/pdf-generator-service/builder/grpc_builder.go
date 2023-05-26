package builder

import (
	"google.golang.org/grpc"
	"net"
	"pdf-generator-service/environment"
	"pdf-generator-service/services/generator"
)

func ServeGRPCServer(services *Services) error {
	listener, err := net.Listen("tcp", environment.GRPC_SERVER_HOST.Value())
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	generator.RegisterGeneratorServiceServer(grpcServer, services.GeneratorService.(*generator.SimpleService))

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
