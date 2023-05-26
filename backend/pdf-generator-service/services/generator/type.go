package generator

import (
	"context"
)

type (
	GRPCService interface {
		GeneratePDFFromTemplateAndData(ctx context.Context, in *GeneratePDFFromTemplateAndDataRequest) (*GeneratePDFFromTemplateAndDataResponse, error)
	}

	Service interface {
		GeneratePDFFromTemplate(template string) ([]byte, error)
		ParseTemplate(templatePath string, data any) (string, error)
		GetAllExamples() ([]DataWithFileName, error)
	}

	SimpleService struct {
		UnimplementedGeneratorServiceServer
	}
)
