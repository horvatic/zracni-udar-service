package model

type Build struct {
	Id           string        `json:"id"`
	ProjectId    string        `json:"project_id"`
	BuildId      string        `json:"build_id"`
	Version      string        `json:"version"`
	Environments []Environment `json:"environments"`
}
