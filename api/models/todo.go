package models

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrTitleEmpty = errors.New("title can't be empty")
	ErrTaskEmpty  = errors.New("task can't be empty")
)

type TODO struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (todo *TODO) Validate() error {
	todo.Title = strings.TrimSpace(todo.Title)
	if todo.Task == "" {
		return ErrTitleEmpty
	}

	todo.Task = strings.TrimSpace(todo.Task)
	if todo.Task == "" {
		return ErrTaskEmpty
	}

	return nil
}
