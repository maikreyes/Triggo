package repository

type InstallationRepositories struct {
	TotalCount   int          `json:"total_count"`
	Repositories []Repository `json:"repositories"`
}
