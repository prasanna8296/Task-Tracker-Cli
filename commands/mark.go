package commands

import (
	"fmt"
	"strconv"
	"task-cli/storage"
	"time"
)

func MarkCommand(args []string, status string) {
	if len(args) < 1 {
		fmt.Println("Usage: mark-<status> <id>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	tasks, _ := storage.LoadTasks()
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			storage.SaveTasks(tasks)
			fmt.Printf("Task marked as %s\n", status)
			return
		}
	}
	fmt.Println("Task not found")
}
