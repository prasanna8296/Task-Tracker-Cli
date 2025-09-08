package commands

import (
    "fmt"
    "strconv"
    "task-cli/storage"
)

func DeleteCommands(args []string) {
    if len(args) < 1 {
        fmt.Println("Usage:delete<id>")
        return
    }
    id, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("invalid id")
        return
    }
    tasks, _ := storage.LoadTasks()
    for i, t := range tasks {
        if t.Id == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            storage.SaveTasks(tasks)
            fmt.Println("Task deleted successfully")
            return
        }
    }
    fmt.Println("Task not found")
}
