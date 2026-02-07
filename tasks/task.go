package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// just taking the description while creating a new task
func New(description string) (*Task, error) {
	// every task must atleast have some description.
	if description == "" {
		return nil, errors.New("A task must have a description !")
	}

	return &Task{
		Id:          uuid.NewString(),
		Description: description,
		Status:      "not done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// retrieve the task that needs to be updated and after retrieving that task
// call this function to update the task.
func (task *Task) Update(description string) error {
	switch description {
	case task.Description:
		return nil
	case "":
		return errors.New("Description cannot be empty...")
	}

	task.Description = description

	// everything all right.
	return nil
}
