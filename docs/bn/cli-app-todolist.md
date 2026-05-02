# প্রজেক্ট: Todo List CLI

এই নোটে [projects/cli-app/todolist/main.go](../../projects/cli-app/todolist/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

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

## এই প্রজেক্টে যা কভার করা হয়েছে

- `os.Args` দিয়ে command-based CLI তৈরি করা
- JSON file-এ task data সংরক্ষণ করা
- Task data model করার জন্য struct ব্যবহার করা
- add, list, done, এবং delete command implement করা
- input validation এবং file error handle করা

## ব্যাখ্যা

1. `os.Args`  
   Command-line arguments নেয়। Program `add`, `list`, `done`, এবং `delete` এর মতো command আশা করে।

2. `type Task struct`  
   প্রতিটি task-এর title, completion status, এবং creation time define করে।

3. `const dataFile = "tasks.json"`  
   Task save এবং load করার জন্য যে file ব্যবহার হবে তার নাম ধরে রাখে।

4. `handleAdd`  
   Argument-গুলো join করে task title বানায়, existing task load করে, নতুন task যোগ করে, তারপর JSON file-এ save করে।

5. `handleList`  
   সব task print করে। Incomplete task-এর জন্য `[ ]` এবং completed task-এর জন্য `[x]` দেখায়।

6. `handleDone`  
   Task number অনুযায়ী task completed হিসেবে mark করে।

7. `handleDelete`  
   Task number অনুযায়ী task delete করে।

8. `loadTasks` এবং `saveTasks`  
   JSON encoding ব্যবহার করে `tasks.json` file read এবং write করে।

9. `parseTaskNumber`  
   Task number positive integer কি না validate করে।

## কিভাবে রান করবে

Project folder থেকে:

```bash
go run main.go list
```

## উদাহরণ command

```bash
go run main.go add "Learn goroutines"
go run main.go list
go run main.go done 1
go run main.go delete 1
```

## উদাহরণ output

```text
1. [ ] Learn goroutines
2. [x] Build a small CLI project
```

## দ্রুত সারসংক্ষেপ

- `os.Args` Go-তে command-line tool তৈরির ভিত্তি।
- JSON local structured data সংরক্ষণের একটি সহজ উপায়।
- ছোট helper function ব্যবহার করলে CLI code পড়তে এবং maintain করতে সুবিধা হয়।
- Input validation common runtime mistake কমায়।