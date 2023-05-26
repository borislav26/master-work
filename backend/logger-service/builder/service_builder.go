package builder

import (
	logger_grpc "logger-service/logs"
	"logger-service/services/logger"
)

type Services struct {
	LoggerService logger.Service
}

func BuildServices(repositories *Repositories) *Services {
	loggerService := &logger.SimpleService{
		UnimplementedLogServiceServer: logger_grpc.UnimplementedLogServiceServer{},
		MongoRepository:               repositories.MongoRepository,
	}
	return &Services{
		LoggerService: loggerService,
	}
}
