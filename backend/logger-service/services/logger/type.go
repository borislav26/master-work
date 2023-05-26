package logger

import (
	"context"
	logger_grpc "logger-service/logs"
	mongo_repo "logger-service/mongo-repo"
)

type (
	Service interface {
		WriteLog(c context.Context, request *logger_grpc.LogRequest) (*logger_grpc.LogResponse, error)
	}

	SimpleService struct {
		logger_grpc.UnimplementedLogServiceServer
		MongoRepository mongo_repo.Repository
	}
)
