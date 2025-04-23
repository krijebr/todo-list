package entity

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}
