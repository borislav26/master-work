package environment

import (
	"os"
)

type Env string

const (
	GRPC_SERVER_HOST Env = "GRPC_SERVER_HOST"
)

func (e Env) Value() string {
	return os.Getenv(string(e))
}
