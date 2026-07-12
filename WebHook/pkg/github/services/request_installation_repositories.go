package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"triggo/pkg/github/model/repository"
)

func (s *Services) RequestInstallationRepositories(Token string) (repository.InstallationRepositories, error) {

	var Repositories repository.InstallationRepositories

	for page := 1; ; page++ {
		requesUrl := fmt.Sprintf("https://api.github.com/installation/repositories?per_page=10&page=%d", page)

		req, err := http.NewRequest(http.MethodGet, requesUrl, nil)

		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		req.Header.Set("Authorization", "Bearer "+Token)

		client := &http.Client{
			Timeout: 20 * time.Second,
		}

		resp, err := client.Do(req)

		if err != nil {
			return repository.InstallationRepositories{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
			bodyBites, _ := io.ReadAll(resp.Body)
			return repository.InstallationRepositories{}, fmt.Errorf("Github reject the request %d: %s", resp.StatusCode, string(bodyBites))
		}

		var tempRepos repository.InstallationRepositories
		err = json.NewDecoder(resp.Body).Decode(&tempRepos)

		if err != nil {

			fmt.Println(err.Error())

			return repository.InstallationRepositories{}, err
		}

		Repositories.TotalCount = tempRepos.TotalCount
		Repositories.Repositories = append(Repositories.Repositories, tempRepos.Repositories...)

		if len(Repositories.Repositories) == Repositories.TotalCount {
			break
		}

	}

	return Repositories, nil

}
