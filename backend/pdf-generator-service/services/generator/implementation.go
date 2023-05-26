package generator

import (
	"bytes"
	"fmt"
	"os"
	error2 "pdf-generator-service/error"
	"strings"
	"text/template"
)

func (s SimpleService) GeneratePDFFromTemplate(template string) ([]byte, error) {
	generatedPdf, err := s.makeRequestToGeneratorService(template)
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}
	return generatedPdf, nil
}

func (s SimpleService) ParseTemplate(templatePath string, data any) (string, error) {
	templ, err := template.ParseFiles(fmt.Sprintf("./templates/%s.html", templatePath))
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err)
	}

	buffer := &bytes.Buffer{}
	err = templ.Execute(buffer, data)
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err)
	}

	err = os.WriteFile(fmt.Sprintf("./temp/%s", "final.html"), buffer.Bytes(), 0644)
	if err != nil {
		return "", error2.ServiceLayerError.Wrap(err)
	}

	return buffer.String(), nil
}

func (s SimpleService) GetAllExamples() ([]DataWithFileName, error) {
	examples := make([]DataWithFileName, 0)
	filesInDirectory, err := os.ReadDir("./examples")
	if err != nil {
		return nil, error2.ServiceLayerError.Wrap(err)
	}
	for _, file := range filesInDirectory {
		f, err := os.ReadFile(fmt.Sprintf("./examples/%s", file.Name()))
		if err != nil {
			return nil, error2.ServiceLayerError.Wrap(err)
		}
		newFile := DataWithFileName{
			FileName: strings.TrimSuffix(file.Name(), ".png"),
			Data:     f,
		}
		examples = append(examples, newFile)
	}
	return examples, nil
}
