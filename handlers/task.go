package handlers

import (
	"errors"

	"github.com/kostovJP/Task-tracker-CLI/fileman"
	"github.com/kostovJP/Task-tracker-CLI/tasks"
	"github.com/kostovJP/Task-tracker-CLI/utils"
)

func RetrieveTaskById(tasks []tasks.Task, taskId int64) (*tasks.Task, error) {
	//task is a copy of each slice element
	//&task points to a loop variable, not the slice element
	//so must return &tasks[index] not &task
	for index, task := range tasks {
		if task.Id == taskId {
			return &tasks[index], nil
		}
	}

	return nil, errors.New("No task exists with the given id..")
}

func AddTask(description, filename string) (int64, error) {
	// creates a new task
	// first check if file exists...
	fileExist := utils.CheckFileExists(filename)

	newTask, err := tasks.New(description)

	if err != nil {
		return 0, err
	}

	var taskList []tasks.Task

	if fileExist {
		taskList, err = fileman.ReadJSON(filename)

		if err != nil {
			return 0, err
		}
	}

	newTask.Id = int64(len(taskList) + 1)

	taskList = append(taskList, *newTask)
	return newTask.Id, fileman.WriteJSON(filename, taskList)
}

func UpdateTask(taskId int64, description, path string) error {
	if description == "" {
		return errors.New("empty description...")
	}

	tasks, err := fileman.ReadJSON(path)

	if err != nil {
		return err
	}

	task, err := RetrieveTaskById(tasks, taskId)

	if err != nil {
		return err
	}

	task.Description = description
	return fileman.WriteJSON(path, tasks)
}

func DeleteTask(taskId int64, path string) error {
	// Deletes task defined by taskId.
	taskList, err := fileman.ReadJSON(path)

	if err != nil {
		return err
	}

	_ , err = RetrieveTaskById(taskList, taskId)

	if err != nil {
		return err
	}

	var newTaskList []tasks.Task

	for _, entry := range taskList {
		if entry.Id != taskId {
			// adjust the ids of the task after the deleted task.
			if entry.Id > taskId {
				entry.Id -= 1
			}
			newTaskList = append(newTaskList, entry)
		}
	}

	return fileman.WriteJSON(path, newTaskList)
}

func MarkTask(taskId int64, newStatus, path string) error {
	taskList, err := fileman.ReadJSON(path)

	if err != nil {
		return err
	}

	task, err := RetrieveTaskById(taskList, taskId)

	if err != nil {
		return err
	}

	task.Status = newStatus
	return fileman.WriteJSON(path, taskList)	
}

// fetches all the tasks and lists them in the terminal
func GetAllTasks(path string) ([]tasks.Task, error) {
	tasks, err := fileman.ReadJSON(path)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
