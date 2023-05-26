package generator

import (
	"context"
)

type (
	GRPCService interface {
		GeneratePDFFromTemplateAndData(context.Context, *GeneratePDFFromTemplateAndDataRequest) (*GeneratePDFFromTemplateAndDataResponse, error)
	}

	SimpleService struct {
	}
)
