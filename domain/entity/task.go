package entity

import (
	"time"
)

type Task struct {
	ID ID
	Title string
	Done bool
	CreatedAt time.Time
	UpdateedAt time.Time
}

func NewTask(title string) Task {
	return Task{
		ID: NewId(),
		Title: title,
		Done: false,
		CreatedAt: time.Now(),
		UpdateedAt: time.Now(),
	}
}