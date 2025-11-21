package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "time"
)

const fileName = "tasks.json"

type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}


func loadTasks() ([]Task, error) {
    data, err := os.ReadFile(fileName)
    if err != nil {
        if os.IsNotExist(err) {
            return []Task{}, nil
        }
        return nil, err
    }

    var tasks []Task
    err = json.Unmarshal(data, &tasks)
    return tasks, err
}

func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644)
}


func addTask(desc string) {
    tasks, _ := loadTasks()

    newTask := Task{
        ID:          len(tasks) + 1,
        Description: desc,
        Status:      "todo",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    tasks = append(tasks, newTask)
    saveTasks(tasks)

    fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func updateTask(id int, desc string) {
    tasks, _ := loadTasks()

    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Description = desc
            tasks[i].UpdatedAt = time.Now()
            saveTasks(tasks)
            fmt.Println("Task updated successfully")
            return
        }
    }

    fmt.Println("Error: Task not found")
}

func deleteTask(id int) {
    tasks, _ := loadTasks()
    newTasks := []Task{}

    found := false
    for _, t := range tasks {
        if t.ID == id {
            found = true
            continue
        }
        newTasks = append(newTasks, t)
    }

    if !found {
        fmt.Println("Error: Task not found")
        return
    }

    saveTasks(newTasks)
    fmt.Println("Task deleted successfully")
}

func setStatus(id int, status string) {
    tasks, _ := loadTasks()

    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Status = status
            tasks[i].UpdatedAt = time.Now()
            saveTasks(tasks)
            fmt.Println("Task status updated successfully")
            return
        }
    }
    fmt.Println("Error: Task not found")
}

func listTasks(filter string) {
    tasks, _ := loadTasks()

    for _, t := range tasks {
        if filter != "" && t.Status != filter {
            continue
        }
        fmt.Printf("%d. [%s] %s (created: %s)\n",
            t.ID, t.Status, t.Description, t.CreatedAt.Format(time.RFC3339))
    }
}


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage:")
        fmt.Println("  task-cli add <description>")
        fmt.Println("  task-cli update <id> <description>")
        fmt.Println("  task-cli delete <id>")
        fmt.Println("  task-cli mark-in-progress <id>")
        fmt.Println("  task-cli mark-done <id>")
        fmt.Println("  task-cli list [status]")
        return
    }

    cmd := os.Args[1]

    switch cmd {

    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Error: Missing description")
            return
        }
        addTask(os.Args[2])

    case "update":
        if len(os.Args) < 4 {
            fmt.Println("Error: Usage: update <id> <description>")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        updateTask(id, os.Args[3])

    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Error: Missing ID")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        deleteTask(id)

    case "mark-in-progress":
        if len(os.Args) < 3 {
            fmt.Println("Error: Missing ID")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        setStatus(id, "in-progress")

    case "mark-done":
        if len(os.Args) < 3 {
            fmt.Println("Error: Missing ID")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        setStatus(id, "done")

    case "list":
        filter := ""
        if len(os.Args) == 3 {
            filter = os.Args[2]
        }
        listTasks(filter)

    default:
        fmt.Println("Unknown command:", cmd)
    }
}
