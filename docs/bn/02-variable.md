# পাঠ ০২: Go-তে Variable

এই নোটে [examples/02-variable/main.go](../../examples/02-variable/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

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

## এই পাঠে যা কভার করা হয়েছে

- `var` দিয়ে explicit variable declaration
- `:=` দিয়ে short declaration
- Go type inference
- ফাংশনের multiple return values
- `_` (blank identifier) দিয়ে অপ্রয়োজনীয় value ignore করা

## ব্যাখ্যা

1. `var name string = "Chitano Kumar"`
   এখানে variable name, type, এবং value সব স্পষ্টভাবে লেখা হয়েছে।

2. `age := 34`
   short declaration syntax। Go স্বয়ংক্রিয়ভাবে `age`-কে `int` হিসেবে infer করে।

3. `email := "chitanokumar@gmail.com"`
   এটিও inferred variable, type হচ্ছে `string`।

4. `count, _ := fmt.Println(email)`
   `fmt.Println` দুইটি value return করে:
   - কত byte লেখা হয়েছে
   - error

   এখানে `count`-এ byte count রাখা হয়েছে, আর `_` দিয়ে error ignore করা হয়েছে।

5. `fmt.Println("Number of character in my email", count)`
   `fmt.Println(email)` থেকে পাওয়া number প্রিন্ট করে।

## উদাহরণ আউটপুট

```text
My name is Chitano Kumar
My age is 34
chitanokumar@gmail.com
Number of character in my email 23
```

## গুরুত্বপূর্ণ নোট

`23` শুধু email-এর character count নয়। `fmt.Println` newline প্রিন্ট করে, সেটিও byte count-এর মধ্যে পড়ে।

শুধু email length পেতে ব্যবহার করতে পারো:

```go
fmt.Println(len(email))
```

## দ্রুত সারসংক্ষেপ

- explicit type চাইলে `var` ব্যবহার করো
- local concise declaration-এর জন্য `:=` ব্যবহার করো
- Go ফাংশন একাধিক value return করতে পারে
- অপ্রয়োজনীয় return value ignore করতে `_` ব্যবহার করো