package logger

import (
	"context"
	"log"
	logger_grpc "logger-service/logs"
)

func (s SimpleService) WriteLog(c context.Context, request *logger_grpc.LogRequest) (*logger_grpc.LogResponse, error) {
	response := &logger_grpc.LogResponse{
		Result: false,
	}

	log.Println("doslo je ovde")

	err := s.MongoRepository.SaveLog(c, request.LogEntry)
	if err != nil {
		return response, err
	}
	response.Result = true
	return response, nil
}
