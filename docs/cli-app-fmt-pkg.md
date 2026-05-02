# Project: CLI App with fmt Package

This note explains the code in [projects/cli-app/fmt-pkg/main.go](../projects/cli-app/fmt-pkg/main.go).

## Source code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Kickstart CLI App")

	//Welcome message
	var name string
	fmt.Println("Enter your name please: ")
	fmt.Scan(&name)
	fmt.Printf("Thenk you, %s for using our CLI app\n", name)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command:")
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, "exit") {
			fmt.Println("Goodbye!")
			return
		}

		fmt.Println("You typed:", text)
	}
}
```

## What this project covers

- Building a simple interactive CLI application
- Reading input with both `fmt.Scan` and `bufio.NewReader`
- Using a loop to keep the program running until the user exits
- Detecting a command with the `strings` package

## Explanation

1. `fmt.Println("Welcome to Kickstart CLI App")`  
   Prints a welcome message when the program starts.

2. `fmt.Scan(&name)`  
   Reads a single word from standard input and stores it in the `name` variable.

3. `bufio.NewReader(os.Stdin)`  
   Creates a buffered reader so the program can read a full line of text from the terminal.

4. `for { ... }`  
   Starts an infinite loop so the CLI can keep accepting commands until the user types `exit`.

5. `reader.ReadString('\n')`  
   Reads input until the Enter key is pressed.

6. `strings.Contains(text, "exit")`  
   Checks whether the typed input contains the word `exit`. If it does, the app prints `Goodbye!` and stops.

7. `fmt.Println("You typed:", text)`  
   Echoes the command back to the user.

## How to run

From the project folder:

```bash
go run main.go
```

## Example session

```text
Welcome to Kickstart CLI App
Enter your name please:
Chitano
Thenk you, Chitano for using our CLI app
Enter command:hello
You typed: hello

Enter command:exit
Goodbye!
```

## Quick recap

- `fmt.Scan` is useful for simple input.
- `bufio.NewReader` is better when you want to read a full line.
- A `for` loop can keep a CLI app running continuously.
- `strings.Contains` helps detect commands like `exit`.