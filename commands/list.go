package commands

import (
	"fmt"
	"strings"
	"task-cli/storage"
)

func ListCommand(args []string) {
	tasks, _ := storage.LoadTasks()
	statusFilter := ""
	if len(args) > 0 {
		statusFilter = strings.ToLower(args[0])
	}

	for _, t := range tasks {
		if statusFilter == "" || t.Status == statusFilter {
			fmt.Printf("Id->[%d] | Description-> %s | Status-> (%s) | CreatedAt-> %s | UpdatedAt-> %s\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
	}
}
