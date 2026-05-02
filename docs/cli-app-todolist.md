# Project: Todo List CLI

This note explains the code in [projects/cli-app/todolist/main.go](../projects/cli-app/todolist/main.go).

## Source code

```go
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
```

## What this project covers

- Building a command-based CLI with `os.Args`
- Saving tasks in a JSON file
- Creating a struct for task data
- Implementing add, list, complete, and delete actions
- Handling input validation and file errors

## Explanation

1. `os.Args`  
   Reads command-line arguments. The program expects commands like `add`, `list`, `done`, and `delete`.

2. `type Task struct`  
   Defines the structure of each task with a title, completion status, and creation time.

3. `const dataFile = "tasks.json"`  
   Stores the file name used to save and load task data.

4. `handleAdd`  
   Joins the remaining arguments into a task title, loads existing tasks, appends the new task, and writes everything back to the JSON file.

5. `handleList`  
   Prints all tasks with `[ ]` for incomplete items and `[x]` for completed ones.

6. `handleDone`  
   Marks a task as completed by its task number.

7. `handleDelete`  
   Removes a task from the list by its task number.

8. `loadTasks` and `saveTasks`  
   Handle reading from and writing to `tasks.json` using JSON encoding.

9. `parseTaskNumber`  
   Validates that the task number is a positive integer.

## How to run

From the project folder:

```bash
go run main.go list
```

## Example commands

```bash
go run main.go add "Learn goroutines"
go run main.go list
go run main.go done 1
go run main.go delete 1
```

## Example output

```text
1. [ ] Learn goroutines
2. [x] Build a small CLI project
```

## Quick recap

- `os.Args` is the base for command-line tools in Go.
- JSON is a simple way to persist structured data locally.
- Small helper functions make CLI code easier to read and maintain.
- Input validation prevents common runtime mistakes.