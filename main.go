package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "task-cli/commands"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        // Show Menu
        fmt.Println("\nTask CLI Menu")
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Update Task")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Mark Task as Done")
        fmt.Println("6. Mark Task as In Progress")
        fmt.Println("7. Exit")

        fmt.Print("\n>> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        choice, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Please enter a valid number.")
            continue
        }

        switch choice {
        case 1:
            fmt.Print("Enter task description: ")
            desc, _ := reader.ReadString('\n')
            desc = strings.TrimSpace(desc)
            if desc == "" {
                fmt.Println("Description cannot be empty.")
                continue
            }
            commands.AddCommands([]string{desc})

        case 2:
            commands.ListCommand([]string{})

        case 3:
            fmt.Print("Enter task ID to update: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)

            fmt.Print("Enter new description: ")
            newDesc, _ := reader.ReadString('\n')
            newDesc = strings.TrimSpace(newDesc)

            commands.UpdateCommand([]string{idStr, newDesc})

        case 4:
            fmt.Print("Enter task ID to delete: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            commands.DeleteCommands([]string{idStr})

        case 5:
            fmt.Print("Enter task ID to mark as done: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            commands.MarkCommand([]string{idStr}, "done")

        case 6:
            fmt.Print("Enter task ID to mark as in-progress: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            commands.MarkCommand([]string{idStr}, "in-progress")

        case 7:
            fmt.Println("Exiting Task CLI.")
            return

        default:
            fmt.Println("Invalid choice. Try again.")
        }
    }
}
