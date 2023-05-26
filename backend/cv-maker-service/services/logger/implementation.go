package logger

import "context"

func (s SimpleService) WriteLog(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, nil
}
