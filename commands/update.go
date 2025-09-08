package commands

import (
	"fmt"
	"strconv"
	"strings"
	"task-cli/storage"
	"time"
)

func UpdateCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: update <id> <new description>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid Id")
		return
	}
	newDesc := strings.Join(args[1:], " ")
	tasks, _ := storage.LoadTasks()

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			storage.SaveTasks(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}

	fmt.Println("Task not found")
}
