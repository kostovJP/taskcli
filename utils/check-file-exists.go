package utils

import (
	"fmt"
	"os"

	"github.com/kostovJP/Task-tracker-CLI/tasks"
)

func CheckFileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

func DisplayTasks(taskList []tasks.Task) {
	for _, task := range taskList {
		fmt.Printf("ID --> %s\nDescription --> %s\nStatus --> %s\n\n---------------\n\n",
			task.Id, task.Description, task.Status)
	}
}
