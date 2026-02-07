package fileman

import (
	"encoding/json"
	"os"

	"github.com/kostovJP/Task-tracker-CLI/tasks"
)

// we will store an array of tasks.
func WriteJSON(filePath string, data interface{}) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	encoding := json.NewEncoder(file)
	err = encoding.Encode(data)

	return err
}

// Note: slices are already like pointers, so when returning slice, 
// no need to return a pointer to a slice ...
func ReadJSON(filePath string) ([]tasks.Task, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var task []tasks.Task

	err = json.Unmarshal(data, &task);

	if err != nil {
		return nil, err
	}

	return task, nil
}
