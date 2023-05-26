package generator

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func (s SimpleService) makeRequestToGeneratorService(file string) ([]byte, error) {
	req, err := http.NewRequest("POST", "http://generator-service:4005/generate-pdf", bytes.NewBuffer([]byte(file)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{
		Timeout: 40 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
