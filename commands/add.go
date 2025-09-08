package commands

import (
	"fmt"
	"strings"
	"task-cli/models"
	"task-cli/storage"
	"time"
)

func AddCommands(args []string) {
	if len(args) < 1 {
		fmt.Println("Task description required")
		return
	}
	description := strings.Join(args, " ")
	tasks, _ := storage.LoadTasks()
	newTask := models.Task{
		Id:          storage.GenNextId(tasks),
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      "todo",
	}
	tasks = append(tasks, newTask)
	storage.SaveTasks(tasks)
	fmt.Printf("Task added successfully (Id: %d)\n", newTask.Id)

}
