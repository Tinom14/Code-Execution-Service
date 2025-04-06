package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/http_service/api/http/types"
)

type HTTPClient struct {
	client    *http.Client
	commitURL string
}

func NewHTTPClient(commitURL string) *HTTPClient {
	return &HTTPClient{
		client:    &http.Client{},
		commitURL: commitURL,
	}
}

func (c *HTTPClient) SendResult(payload types.PostTaskCommitRequest) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("result payload: %w", err)
	}

	resp, err := c.client.Post(c.commitURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return fmt.Errorf("sending result to HTTP server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("HTTP server returned non-OK status: %s", resp.Status)
	}

	return nil
}
