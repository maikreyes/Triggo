package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"triggo/pkg/discord/model/payload"
)

func (s *Services) SendPayload(p payload.Payload) error {

	b, err := json.Marshal(p)

	if err != nil {
		return fmt.Errorf("error converting payload to JSON: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, s.Config.DUrl, bytes.NewBuffer(b))

	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error sending request to Discord: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		bodyBites, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("discord rejected the request, status %d: %s", resp.StatusCode, string(bodyBites))
	}

	return nil
}
