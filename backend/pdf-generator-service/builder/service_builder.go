package builder

import (
	"pdf-generator-service/services/authentication"
	"pdf-generator-service/services/generator"
)

type Services struct {
	GeneratorService          generator.Service
	AuthenticationGRPCService authentication.GRPCService
}

func BuildServices() *Services {
	generatorService := &generator.SimpleService{
		UnimplementedGeneratorServiceServer: generator.UnimplementedGeneratorServiceServer{},
	}

	authenticationService := &authentication.SimpleService{}

	return &Services{
		GeneratorService:          generatorService,
		AuthenticationGRPCService: authenticationService,
	}
}
