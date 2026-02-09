package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/kostovJP/Task-tracker-CLI/handlers"
	"github.com/kostovJP/Task-tracker-CLI/utils"
)

func main() {
	// run a switch case type thing, first find out the second string in the
	// command line argument and see what it is,
	// based on that decide what operations to perform, i.e add, update, delete
	if len(os.Args) < 2 || len(os.Args) > 4 {
		panic("incorrect number of arguments (expected min 2 max 4)")
	}

	args := os.Args[1:]
	command := args[0]
	filePath := "tasks.json"

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("incorrect number of arguments !! (expected 2)")
			fmt.Println("use: taskcli add <task-description>")
			return
		}

		taskDescription := args[1]
		taskId, err := handlers.AddTask(taskDescription, filePath)
		if err != nil {
			fmt.Println(err)
			panic("couldnot add task!!")
		}
		fmt.Printf("Task added successfully (ID: %d)", taskId)

	case "update":
		if len(args) < 2 {
			fmt.Println("incorrect number of arguments !! (expected 2)")
			fmt.Println("use: taskcli update <task-id>")
			return
		}

		taskDescription := args[2]
		taskId, err := strconv.ParseInt(args[1], 10, 64)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}
		err = handlers.UpdateTask(taskId, taskDescription, filePath)
		if err != nil {
			fmt.Println(err)
			panic("couldnot update task!!")
		}
		fmt.Printf("Task %d updated successfully!", taskId)

	case "delete":
		if len(args) < 2 {
			fmt.Println("incorrect number of arguments !! (expected 2)")
			fmt.Println("use: taskcli delete <task-id>")
			return
		}

		taskId, err := strconv.ParseInt(args[1], 10, 64)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}
		err = handlers.DeleteTask(taskId, filePath)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}

	case "mark-in-progress":
		newStatus := "in-progress"
		taskId, err := strconv.ParseInt(args[1], 10, 64)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}

		err = handlers.MarkTask(taskId, newStatus, filePath)

		if err != nil {
			fmt.Println(err)
			panic("unable to change file status to in-progress")
		}

	case "mark-done":
		newStatus := "done"
		taskId, err := strconv.ParseInt(args[1], 10, 64)

		if err != nil {
			fmt.Println(err)
			panic("unable to delete task!!")
		}

		err = handlers.MarkTask(taskId, newStatus, filePath)
		if err != nil {
			fmt.Println(err)
			panic("unable to change file status to in-progress")
		}

	case "list":
		var fileStatus string
		if len(args) == 1 {
			fileStatus = ""
		} else {
			fileStatus = args[1]
		}

		if (fileStatus != "in-progress") && (fileStatus != "done") && (fileStatus != "") {
			fmt.Println("No such status as: ", fileStatus)
			fmt.Println(`Available status are: "in-progress", "done"`)
			return
		}

		taskList, err := handlers.GetAllTasks(filePath, fileStatus)

		if err != nil {
			fmt.Println("No tasks.json file found in current directory...")
			fmt.Println(`Create one by using: $taskcli add "task description"`)
			return
		}

		utils.DisplayTasks(taskList)
	case "--help":
		utils.ShowHelp() // show the help menu.
	default:
		fmt.Println(command, ": No such command found..")
	}
}
