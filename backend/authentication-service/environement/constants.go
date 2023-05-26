package environement

import "os"

type Env string

const (
	GRPC_SERVER_HOST         Env = "GRPC_SERVER_HOST"
	HTTP_SERVER_PORT         Env = "HTTP_SERVER_PORT"
	DATABASE_URL             Env = "DATABASE_URL"
	JWT_SECRET_TOKEN_KEY     Env = "JWT_SECRET_TOKEN_KEY"
	GRPC_LOGGER_SERVICE_HOST Env = "GRPC_LOGGER_SERVICE_HOST"
	SERVICE_NAME             Env = "SERVICE_NAME"
)

func (e Env) String() string {
	return string(e)
}

func (e Env) GetValue() string {
	return os.Getenv(e.String())
}
