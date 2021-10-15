package model

type ServiceInfo struct {
	Id          string `json:"id"`
	ProjectId   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	HealthUri   string `json:"health_uri"`
}
