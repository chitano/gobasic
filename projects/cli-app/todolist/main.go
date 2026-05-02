package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const dataFile = "tasks.json"

type Task struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "add":
		handleAdd(os.Args[2:])
	case "list":
		handleList()
	case "done":
		handleDone(os.Args[2:])
	case "delete":
		handleDelete(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
	}
}

func handleAdd(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: todo add \"Task title\"")
		return
	}

	title := strings.TrimSpace(strings.Join(args, " "))
	if title == "" {
		fmt.Println("Task title cannot be empty.")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Failed to load tasks: %v\n", err)
		return
	}

	tasks = append(tasks, Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	})

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Failed to save tasks: %v\n", err)
		return
	}

	fmt.Printf("Added task: %s\n", title)
}

func handleList() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Failed to load tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Title)
	}
}

func handleDone(args []string) {
	tasks, index, err := loadTasksAndIndex(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if tasks[index].Completed {
		fmt.Printf("Task %d is already completed.\n", index+1)
		return
	}

	tasks[index].Completed = true

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Failed to save tasks: %v\n", err)
		return
	}

	fmt.Printf("Marked task %d as complete.\n", index+1)
}

func handleDelete(args []string) {
	tasks, index, err := loadTasksAndIndex(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	deletedTitle := tasks[index].Title
	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Failed to save tasks: %v\n", err)
		return
	}

	fmt.Printf("Deleted task %d: %s\n", index+1, deletedTitle)
}

func loadTasksAndIndex(args []string) ([]Task, int, error) {
	if len(args) != 1 {
		return nil, 0, errors.New("Please provide exactly one task number.")
	}

	index, err := parseTaskNumber(args[0])
	if err != nil {
		return nil, 0, err
	}

	tasks, err := loadTasks()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to load tasks: %w", err)
	}

	if len(tasks) == 0 {
		return nil, 0, errors.New("No tasks found.")
	}

	if index < 0 || index >= len(tasks) {
		return nil, 0, fmt.Errorf("task number must be between 1 and %d", len(tasks))
	}

	return tasks, index, nil
}

func parseTaskNumber(raw string) (int, error) {
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errors.New("Task number must be a valid integer.")
	}
	if n <= 0 {
		return 0, errors.New("Task number must be greater than 0.")
	}
	return n - 1, nil
}

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	if len(strings.TrimSpace(string(data))) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0o644)
}

func printUsage() {
	fmt.Println("Todo List CLI")
	fmt.Println("Usage:")
	fmt.Println("  todo add \"Task title\"")
	fmt.Println("  todo list")
	fmt.Println("  todo done <task-number>")
	fmt.Println("  todo delete <task-number>")
}
