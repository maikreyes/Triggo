package push

import "triggo/internal/github/model/commit"

type GithubPush struct {
	Commit     commit.GithubCommit `json:"head_commit"`
	Ref        string              `json:"ref"`
	Repository string              `json:"repository"`
	Sender     string              `json:"sender"`
}
