package entity

import (
	"time"
)

type Task struct {
	ID int
	Title string
	Done bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewTask(id int, title string) Task {
	return Task{
		ID: id,
		Title: title,
		Done: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}