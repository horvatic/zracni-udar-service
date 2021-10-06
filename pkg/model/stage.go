package model

type Stage struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Order  int    `json:"order"`
}
