package storage

import (
	"encoding/json"
	"errors"
	"os"
	"task-cli/models"
)

const taskFile = "tasks.json"

func LoadTasks() ([]models.Task, error) {

	if _, err := os.Stat(taskFile); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(taskFile, []byte("[]"), 0o644)
	}
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "")

	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)

}
func GenNextId(tasks []models.Task) int {
	maxId := 0
	for _, t := range tasks {

		if t.Id > maxId {
			maxId = t.Id
		}
	}
	return maxId + 1
}
