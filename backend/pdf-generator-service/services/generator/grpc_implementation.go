package generator

import (
	"context"
	"log"
)

func (s SimpleService) GeneratePDFFromTemplateAndData(ctx context.Context, in *GeneratePDFFromTemplateAndDataRequest) (*GeneratePDFFromTemplateAndDataResponse, error) {
	res := &GeneratePDFFromTemplateAndDataResponse{
		Result: false,
	}

	log.Println("ovo je template")

	templateBytes, err := s.ParseTemplate(in.Template, in.UserData)
	if err != nil {
		return res, err
	}

	log.Println("ovo je template", templateBytes)
	cvBytes, err := s.GeneratePDFFromTemplate(templateBytes)
	if err != nil {
		return res, err
	}

	res.Result = true
	res.PdfBytes = cvBytes

	return res, nil
}
