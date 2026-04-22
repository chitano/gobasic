# Lesson 07: Custom Data Types in Go

This note explains the code in [examples/07-custome-data-type/main.go](../examples/07-custome-data-type/main.go).

## Source code

```go
package main

import (
	"fmt"
	"regexp"
	"time"
)

type userID uint32
type Email string
type AddressId uint32
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Address struct {
	Id           AddressId
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	Id      userID
	Name    string
	Email   Email
	Address Address
	Phone   string
	Gender  Gender
}

func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	}
	return false
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (e Email) IsValid() bool {
	return emailRegex.MatchString(string(e))
}

func main() {
	now := time.Now()

	users := []User{
		{Id: 1, Name: "Alice Johnson", Email: "alice.johnsonexample.com", Gender: Female, ...},
		{Id: 2, Name: "Bob Smith",     Email: "bob.smith@example.com",    Gender: Male,   ...},
		// ... 5 users total
	}

	for _, u := range users {
		fmt.Printf("User: %-15s | Email: %-30s | Valid: %v\n", u.Name, u.Email, u.Email.IsValid())
	}
}
```

## What this lesson covers

- Defining custom named types based on primitives (`uint32`, `string`)
- Typed constants with a custom `Gender` type
- Defining structs (`Address`, `User`) with multiple fields
- Attaching methods to custom types (method receivers)
- Email validation using `regexp`
- Slices of structs
- Formatted output with `fmt.Printf`

## Explanation

### Custom types

```go
type userID    uint32
type Email     string
type AddressId uint32
type Gender    string
```

These are new distinct types based on existing primitives. Even though `Email` is backed by `string`, you cannot accidentally assign a plain `string` to an `Email` variable without an explicit conversion. This adds type safety.

### Typed constants

```go
const (
	Male   Gender = "male"
	Female Gender = "female"
)
```

Instead of using raw strings like `"male"` everywhere, you define named constants of type `Gender`. This makes the code self-documenting and prevents typos.

### Structs

```go
type Address struct { ... }
type User struct { ... }
```

A struct groups related fields together under one name. `User` embeds an `Address` directly as a field (composition, not inheritance).

### Method on `Gender` — `IsValid()`

```go
func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	}
	return false
}
```

A method is a function with a receiver. Here `g` is a value receiver of type `Gender`. The method returns `true` only if the gender is one of the defined constants.

### Email validation — `IsValid()`

```go
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (e Email) IsValid() bool {
	return emailRegex.MatchString(string(e))
}
```

- `regexp.MustCompile` compiles the regex once at package level (efficient — no recompilation on every call).
- The method converts the `Email` type to a plain `string` before matching.
- The regex enforces the `local@domain.tld` pattern.

### Slice of structs

```go
users := []User{ ... }
```

Creates a slice (dynamic list) of `User` structs with 5 pre-populated demo records.

### Iterating and printing

```go
for _, u := range users {
	fmt.Printf("User: %-15s | Email: %-30s | Valid: %v\n", u.Name, u.Email, u.Email.IsValid())
}
```

- `range` iterates over each element in the slice.
- `%-15s` left-aligns the string in a 15-character-wide column.
- `%v` prints any value in its default format (`true`/`false` for bool).

## Output from this example

```text
User: Alice Johnson    | Email: alice.johnsonexample.com    | Valid: false
User: Bob Smith        | Email: bob.smith@example.com       | Valid: true
User: Carol Lee        | Email: carol.lee@example.com       | Valid: true
User: David Brown      | Email: david.brown@example.com     | Valid: true
User: Eva Green        | Email: eva.green@example.com       | Valid: true
```

> Alice's email is missing the `@` symbol, so `IsValid()` correctly returns `false`.

## Important notes

- Custom types based on primitives are **not interchangeable** without explicit conversion: `string(myEmail)`.
- `regexp.MustCompile` panics at startup if the regex is invalid — use it only for compile-time-known patterns.
- Method receivers in Go can be a value receiver `(t Type)` or a pointer receiver `(t *Type)`. Use pointer receivers when you need to mutate the value.

## Quick recap

- Use `type` to create named types for more expressive and type-safe code.
- Attach behaviour to types with methods.
- Use `regexp` for pattern-based validation.
- Structs + methods in Go are the foundation of object-oriented-style design.
