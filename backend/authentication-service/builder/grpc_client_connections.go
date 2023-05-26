package builder

import (
	"authentication-service/environement"
	logger_grpc "authentication-service/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCClientConnections struct {
	LoggerClientConnection logger_grpc.LogServiceClient
}

type GRPCClients struct {
	LoggerClient *grpc.ClientConn
}

func (g *GRPCClients) GetLoggerClient() error {
	var err error
	g.LoggerClient, err = grpc.Dial(environement.GRPC_LOGGER_SERVICE_HOST.GetValue(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	return err
}

func (g GRPCClients) BuildGRPCServices() *GRPCClientConnections {
	logger := logger_grpc.NewLogServiceClient(g.LoggerClient)

	return &GRPCClientConnections{
		LoggerClientConnection: logger,
	}
}

func NewGRPCClientConnections() *GRPCClientConnections {
	connections := &GRPCClients{}
	err := connections.GetLoggerClient()
	if err != nil {
		log.Fatalln("Failed to get logger client")
	}
	return connections.BuildGRPCServices()
}
