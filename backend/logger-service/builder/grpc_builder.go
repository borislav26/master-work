package builder

import (
	"google.golang.org/grpc"
	"logger-service/environment"
	logger_grpc "logger-service/logs"
	"logger-service/services/logger"
	"net"
)

func ServeGRPCServer(services *Services) error {
	listener, err := net.Listen("tcp", environment.GRPC_SERVER_HOST.Value())
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	logger_grpc.RegisterLogServiceServer(grpcServer, services.LoggerService.(*logger.SimpleService))

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
