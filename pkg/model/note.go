package model

type Note struct {
	Id        string `json:"id"`
	ProjectId string `json:"project_id"`
	Name      string `json:"name"`
	Note      string `json:"note"`
}
