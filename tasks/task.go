package tasks

import (
	"errors"
	"time"
)

type Task struct {
	Id          int64 `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// just taking the description while creating a new task
func New(description string) (*Task, error) {
	// every task must atleast have some description.
	if description == "" {
		return nil, errors.New("A task must have a description !")
	}

	return &Task{
		Id:          0,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

