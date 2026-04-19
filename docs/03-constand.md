# Lesson 03: Constants in Go

This note explains the code in [examples/constant/main.go](../examples/constant/main.go).

## Source code

```go
package main

import "fmt"

const (
	Hostname string = "http://localhost" //This is exportable constant we can use it anywhere in our application but coudn't change it value.
	DBHost   string = "localhost:3036"
	AppName  string = "Go learning"
)

func main() {
	fmt.Println(Hostname)
	//Hostname = "Value not replaceable" //As this is a constant we can change it value here.
	Hostname := "mydomain.com"      // This is another variable with same name
	Hostname = "Now it replaceable" //Now we can change variable value here.
	fmt.Print(Hostname)
}
```

## What this lesson covers

- Declaring constants with `const`
- Grouping multiple constants in a const block
- Difference between constants and variables
- Variable shadowing with `:=`

## Explanation

1. `const (...)`
	You declare three package-level constants: `Hostname`, `DBHost`, and `AppName`.

2. `Hostname` starts with a capital letter
	In Go, that makes it exported (visible outside the package), if used in a non-`main` package.

3. `fmt.Println(Hostname)`
	Prints the constant value `http://localhost`.

4. `//Hostname = "Value not replaceable"`
	This line is commented because it would fail if uncommented.
	Constants cannot be reassigned.

5. `Hostname := "mydomain.com"`
	This creates a new local variable named `Hostname` inside `main`.
	It shadows (hides) the constant with the same name in this scope.

6. `Hostname = "Now it replaceable"`
	This works because now `Hostname` refers to the local variable, not the constant.

7. `fmt.Print(Hostname)`
	Prints the updated local variable value.

## Output from this example

```text
http://localhost
Now it replaceable
```

## Important notes

- A constant value cannot be changed after declaration.
- Shadowing is allowed, but using the same name for constants and local variables can reduce readability.
- `DBHost` and `AppName` are declared but not used in `main` yet.

## Quick recap

- Use `const` for fixed values.
- Use `:=` to declare local variables.
- A local variable can shadow a package-level constant with the same name.
