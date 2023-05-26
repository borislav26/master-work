package environment

import (
	"os"
)

type Env string

const (
	MONGO_DATABASE_URL              Env = "MONGO_DATABASE_URL"
	GRPC_SERVER_HOST                Env = "GRPC_SERVER_HOST"
	GRPC_PDF_GENERATOR_SERVICE_HOST Env = "GRPC_PDF_GENERATOR_SERVICE_HOST"
	GRPC_LOGGER_SERVICE_HOST        Env = "GRPC_LOGGER_SERVICE_HOST"
)

func (e Env) Value() string {
	return os.Getenv(string(e))
}
