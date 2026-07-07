package commit

type GithubCommit struct {
	Committer GithubCommitter `json:"committer"`
	Distinct  bool            `json:"distinct"`
	Id        string          `json:"id"`
	Message   string          `json:"message"`
}
