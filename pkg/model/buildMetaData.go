package model

type BuildMetaData struct {
	Id          string `json:"id"`
	ProjectId   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	RepoOwner   string `json:"repo_owner"`
	RepoName    string `json:"repo_name"`
}
