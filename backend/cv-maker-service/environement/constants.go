package environement

import "os"

type Env string

const (
	GRPC_SERVER_HOST                 Env = "GRPC_SERVER_HOST"
	HTTP_SERVER_PORT                 Env = "HTTP_SERVER_PORT"
	DATABASE_URL                     Env = "DATABASE_URL"
	JWT_SECRET_TOKEN_KEY             Env = "JWT_SECRET_TOKEN_KEY"
	GRPC_AUTHENTICATION_SERVICE_HOST Env = "GRPC_AUTHENTICATION_SERVICE_HOST"
	GRPC_PDF_GENERATOR_SERVICE_HOST  Env = "GRPC_PDF_GENERATOR_SERVICE_HOST"
	GRPC_MAILER_SERVICE_HOST         Env = "GRPC_MAILER_SERVICE_HOST"
)

func (e Env) String() string {
	return string(e)
}

func (e Env) GetValue() string {
	return os.Getenv(e.String())
}
