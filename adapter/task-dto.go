package adapter

type TaskDTO struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}