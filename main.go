package main

import (
	"fmt"
	"os"

	"github.com/kostovJP/Task-tracker-CLI/handlers"
	"github.com/kostovJP/Task-tracker-CLI/utils"
)

func main() {
	// run a switch case type thing, first find out the second string in the
	// command line argument and see what it is,
	// based on that decide what operations to perform, i.e add, update, delete
	args := os.Args[1:]
	command := args[0]
	filePath := "tasks.json"

	switch command {
	case "add":
		taskDescription := args[1]
		taskId, err := handlers.AddTask(taskDescription, filePath)
		if err != nil {
			fmt.Println(err)
			panic("couldnot add task!!")
		}
		fmt.Printf("Task added successfully (ID: %s)", taskId)

	case "update":
		taskDescription := args[2]
		taskId := args[1]
		err := handlers.UpdateTask(taskId, taskDescription, filePath)
		if err != nil {
			fmt.Println(err)
			panic("couldnot update task!!")
		}
		fmt.Printf("Task %s updated successfully!", taskId)

	case "delete":
		taskId := args[1]
		err := handlers.DeleteTask(taskId, filePath)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}

	case "mark-in-progress":
		newStatus := "in-progress"
		taskId := args[1]

		err := handlers.MarkTask(taskId, newStatus, filePath)
		if err != nil {
			fmt.Println(err)
			panic("unable to change file status to in-progress")
		}

	case "mark-done":
		newStatus := "done"
		taskId := args[1]

		err := handlers.MarkTask(taskId, newStatus, filePath)
		if err != nil {
			fmt.Println(err)
			panic("unable to change file status to in-progress")
		}

	case "list":
		taskList, err := handlers.GetAllTasks(filePath)

		if err != nil {
			fmt.Println(err)
			panic("unable to display tasks...")
		}

		utils.DisplayTasks(taskList)
	default:
		fmt.Println(command, ": No such command found..")
	}
}
