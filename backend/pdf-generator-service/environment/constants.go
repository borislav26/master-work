package environment

import "os"

type Env string

const (
	APP_ROOT                        Env = "APP_ROOT"
	GRPC_SERVER_HOST                Env = "GRPC_SERVER_HOST"
	HTTP_SERVER_HOST                Env = "HTTP_SERVER_HOST"
	GRPC_PDF_GENERATOR_SERVICE_HOST Env = "GRPC_PDF_GENERATOR_SERVICE_HOST"
	GRPC_LOGGER_SERVICE_HOST        Env = "GRPC_LOGGER_SERVICE_HOST"
)

func (e Env) String() string {
	return string(e)
}

func (e Env) Value() string {
	return os.Getenv(e.String())
}
