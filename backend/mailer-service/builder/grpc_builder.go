package builder

import (
	"google.golang.org/grpc"
	"mailer-service/environment"
	"mailer-service/services/mailer"
	"net"
)

func ServeGRPCServer(services *Services) error {
	listener, err := net.Listen("tcp", environment.GRPC_SERVER_HOST.Value())
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	mailer.RegisterMailerServiceServer(grpcServer, services.MailerService.(*mailer.SimpleService))

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
