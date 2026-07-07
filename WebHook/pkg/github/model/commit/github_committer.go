package commit

type GithubCommitter struct {
	Date     string `json:"date"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
