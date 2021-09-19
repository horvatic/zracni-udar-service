package model

type Video struct {
	Id          string `json:"id"`
	ProjectId   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
}
