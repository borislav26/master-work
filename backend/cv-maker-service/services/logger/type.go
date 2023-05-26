package logger

import "context"

type (
	GRPCService interface {
		WriteLog(ctx context.Context, req *LogRequest) (*LogResponse, error)
	}

	SimpleService struct {
		UnimplementedLogServiceServer
	}
)
