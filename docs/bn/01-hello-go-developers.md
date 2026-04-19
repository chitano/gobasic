# পাঠ ০১: Hello Go Developers

এই নোটে [examples/01-hello-go-developers/main.go](../../examples/01-hello-go-developers/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

## এই পাঠে যা আছে

- Go প্রোগ্রামের বেসিক স্ট্রাকচার
- `package main` এর ব্যবহার
- `import` দিয়ে প্যাকেজ আনা
- `main()` ফাংশন কেন entry point
- `fmt.Println` দিয়ে আউটপুট প্রিন্ট করা

## ব্যাখ্যা

1. `package main`
   executable প্রোগ্রাম বানাতে Go-তে `main` প্যাকেজ লাগে।

2. `import "fmt"`
   `fmt` প্যাকেজ থেকে formatted input/output ফাংশন ব্যবহার করা যায়।

3. `func main()`
   এটি প্রোগ্রামের শুরু হওয়ার জায়গা। Go runtime এখান থেকেই execute শুরু করে।

4. `fmt.Println("Hello World")`
   কনসোলে `Hello World` প্রিন্ট করে এবং শেষে newline যোগ করে।

## আউটপুট

```text
Hello World
```

## দ্রুত সারসংক্ষেপ

- `main` প্যাকেজ + `main()` ফাংশন = runnable Go program
- `fmt.Println` নতুনদের জন্য সবচেয়ে সহজ output method