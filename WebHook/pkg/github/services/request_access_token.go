package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"triggo/pkg/github/model/installation"
)

func (s *Services) RequestAccessToken(installationID string) (installation.InstallationAccesToken, error) {

	var AccesToken installation.InstallationAccesToken

	requestUrl := "https://api.github.com/app/installations/" + installationID + "/access_tokens"

	webhookJWT, err := s.JWTServices.CreateJWT()

	if err != nil {
		return installation.InstallationAccesToken{}, err
	}

	req, err := http.NewRequest(http.MethodPost, requestUrl, nil)

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+webhookJWT)

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return installation.InstallationAccesToken{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		bodyBites, _ := io.ReadAll(resp.Body)
		return installation.InstallationAccesToken{}, fmt.Errorf("Github reject the request %d: %s", resp.StatusCode, string(bodyBites))
	}

	err = json.NewDecoder(resp.Body).Decode(&AccesToken)

	if err != nil {
		return installation.InstallationAccesToken{}, err
	}

	return AccesToken, nil

}
