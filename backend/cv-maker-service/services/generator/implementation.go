package generator

import (
	"context"
	"cv-maker-service/environement"
	error2 "cv-maker-service/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s SimpleService) GeneratePDFFromTemplateAndData(ctx context.Context, req *GeneratePDFFromTemplateAndDataRequest) (*GeneratePDFFromTemplateAndDataResponse, error) {
	conn, err := grpc.Dial(environement.GRPC_PDF_GENERATOR_SERVICE_HOST.GetValue(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}

	generatorClient := NewGeneratorServiceClient(conn)
	res, err := generatorClient.GeneratePDFFromTemplateAndData(ctx, req)
	if err != nil {
		return res, error2.ServiceLayerError.Wrap(err)
	}

	if !res.Result {
		return res, error2.ServiceLayerError.Wrap(err)
	}

	return res, nil
}
