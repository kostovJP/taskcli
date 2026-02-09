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
		fmt.Printf("_______________________\nID --> %d\nDescription --> %s\nStatus --> %s\n\n",
			task.Id, task.Description, task.Status)
	}
}

func ShowHelp() {
	fmt.Println(`
	-------------------------------------------------------
	                  WELCOME TO taskcli                   
	-------------------------------------------------------
	Adding a new task: (generates a task.json file automatically)
	> task-cli add "Buy groceries"
	___________________________________________________________
	Updating a task:
	> task-cli update <task-Id> "Buy groceries and cook dinner"
	___________________________________________________________
	Deleting a task:
	> task-cli delete <task-Id>
	_______________________________________________________
	Marking a task as in progress or done
	> task-cli mark-in-progress <task-id>
	> task-cli mark-done <task-id>
	_______________________________________________________
	Listing all tasks
	> task-cli list
	_______________________________________________________
	Listing tasks by status
	> task-cli list done
    > task-cli list todo
    > task-cli list in-progress
	_______________________________________________________`)
}