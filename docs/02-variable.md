# Lesson 02: Variables in Go

This note explains the code in [examples/02-variable/main.go](examples/02-variable/main.go).

## Source code

```go
package main

import "fmt"

func main() {
	var name string = "Chitano Kumar" // Explicitly define type.
	fmt.Println("My name is", name)

	age := 34 // Auto-detected type (type inference)
	fmt.Println("My age is", age)

	email := "chitanokumar@gmail.com"

	count, _ := fmt.Println(email)

	fmt.Println("Number of character in my email", count)
}
```

## What this lesson covers

- Explicit variable declaration with `var`
- Short variable declaration with `:=`
- Go type inference
- Multiple return values from a function
- Using `_` (blank identifier) to ignore a value

## Explanation

1. `var name string = "Chitano Kumar"`
	This is explicit declaration. You define the variable name, type, and value.

2. `age := 34`
	This uses short declaration syntax. Go infers `age` as an `int`.

3. `email := "chitanokumar@gmail.com"`
	Another inferred variable, this time a `string`.

4. `count, _ := fmt.Println(email)`
	`fmt.Println` returns two values:
	- number of bytes written
	- error

	Here, `count` stores bytes written, and `_` ignores the error value.

5. `fmt.Println("Number of character in my email", count)`
	Prints the number returned by `fmt.Println(email)`.

## Output from this example

```text
My name is Chitano Kumar
My age is 34
chitanokumar@gmail.com
Number of character in my email 23
```

## Important note

The value `23` is not only the email character count. It includes the newline printed by `fmt.Println`.

If you want the exact email length, use:

```go
fmt.Println(len(email))
```

## Quick recap

- Use `var` when you want an explicit type
- Use `:=` for concise local declarations
- Go functions can return multiple values
- Use `_` when you intentionally do not need one returned value
