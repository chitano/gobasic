# Lesson 04: The fmt Package in Go

This note explains the code in [examples/04-fmt-package/main.go](../examples/04-fmt-package/main.go).

## Source code

```go
package main

import "fmt"

func main() {
	// In most programming language we have built in packages for showing output to screen, console or any other channel. And olso a way to take input from users.
	// In go we have "fmt" package for standard input/output.
	//Simple output, you can just print/output anything.
	fmt.Print("Hello Go Developers")

	//formatted output, you can give a structure in your output.
	name := "Chitano Kumar"
	age := 32
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	//If you want a new line after your output then you can use it. Also it will add  space between all argruments.
	fmt.Println("My name is", name, "and I am ", age, "years old")
}
```

## What this lesson covers

- Importing and using the `fmt` package
- Difference between `fmt.Print`, `fmt.Printf`, and `fmt.Println`
- Using format verbs (`%s`, `%d`) to build formatted strings
- Automatic spacing and newlines

## Explanation

1. `import "fmt"`  
   Imports Go's built-in formatting and I/O package. It is used for printing to the console and reading user input.

2. `fmt.Print("Hello Go Developers")`  
   Prints the string as-is — no newline is added at the end, and no spaces are inserted between arguments.

3. `fmt.Printf("My name is %s and I am %d years old\n", name, age)`  
   Prints a formatted string. Format verbs replace variables:
   - `%s` → string value (`name`)
   - `%d` → integer value (`age`)
   - `\n` → explicit newline character

4. `fmt.Println("My name is", name, "and I am ", age, "years old")`  
   Prints all arguments separated by spaces and adds a newline automatically at the end.

## Common format verbs

| Verb  | Type    | Description                        |
|-------|---------|------------------------------------|
| `%s`  | string  | Plain string                       |
| `%d`  | integer | Decimal integer                    |
| `%f`  | float   | Floating-point number              |
| `%v`  | any     | Default format for any value       |
| `%+v` | struct  | Struct with field names            |
| `%T`  | any     | Type of the value                  |
| `%t`  | bool    | Boolean (`true` or `false`)        |
| `\n`  | —       | Newline (used inside `Printf`)     |

## Output from this example

```text
Hello Go DevelopersMy name is Chitano Kumar and I am 32 years old
My name is Chitano Kumar and I am  32 years old
```

> Note: `fmt.Print` does not add a newline, so the `Printf` output starts on the same line.

## Quick recap

- Use `fmt.Print` when you want raw output with no extra newline.
- Use `fmt.Printf` when you need formatted output with verbs like `%s` and `%d`.
- Use `fmt.Println` for the simplest output — it adds spaces between arguments and a newline at the end.
