package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"triggo/pkg/discord/model/payload"
	messainfromation "triggo/pkg/github/model/messa_infromation"
)

func (s *Services) SendPayload(p payload.Payload, f messainfromation.MessaInformation) error {

	b, err := json.Marshal(p)

	if err != nil {
		return fmt.Errorf("error converting payload to JSON: %w", err)
	}

	installationId, err := strconv.ParseInt(f.Installation.Id, 10, 64)

	if err != nil {
		return fmt.Errorf("Installation Id don´t exist")
	}

	userInformation, err := s.Repository.SearchRecord(installationId, f.Repository.Name)

	if err != nil {
		return fmt.Errorf("User don´t found")
	}

	req, err := http.NewRequest(http.MethodPost, userInformation.DiscordUrl, bytes.NewBuffer(b))

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
