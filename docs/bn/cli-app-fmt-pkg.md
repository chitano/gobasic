# প্রজেক্ট: fmt প্যাকেজ দিয়ে CLI অ্যাপ

এই নোটে [projects/cli-app/fmt-pkg/main.go](../../projects/cli-app/fmt-pkg/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

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

## এই প্রজেক্টে যা কভার করা হয়েছে

- একটি simple interactive CLI app তৈরি করা
- `fmt.Scan` এবং `bufio.NewReader` দিয়ে input নেওয়া
- loop ব্যবহার করে program চালু রাখা
- `strings` package দিয়ে command detect করা

## ব্যাখ্যা

1. `fmt.Println("Welcome to Kickstart CLI App")`  
   Program শুরু হলে একটি welcome message print করে।

2. `fmt.Scan(&name)`  
   Standard input থেকে একটি single word পড়ে `name` variable-এ রাখে।

3. `bufio.NewReader(os.Stdin)`  
   Buffered reader তৈরি করে, যাতে terminal থেকে full line input পড়া যায়।

4. `for { ... }`  
   Infinite loop চালায়, যাতে user `exit` না দেওয়া পর্যন্ত CLI app command নিতে থাকে।

5. `reader.ReadString('\n')`  
   Enter চাপা পর্যন্ত input পড়ে।

6. `strings.Contains(text, "exit")`  
   Input-এর মধ্যে `exit` আছে কি না check করে। থাকলে app `Goodbye!` print করে বন্ধ হয়ে যায়।

7. `fmt.Println("You typed:", text)`  
   User কী লিখেছে তা আবার দেখায়।

## কিভাবে রান করবে

Project folder থেকে:

```bash
go run main.go
```

## উদাহরণ সেশন

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

## দ্রুত সারসংক্ষেপ

- `fmt.Scan` simple input নেওয়ার জন্য উপকারী।
- Full line input নেওয়ার জন্য `bufio.NewReader` বেশি সুবিধাজনক।
- `for` loop দিয়ে CLI app বারবার command নিতে পারে।
- `strings.Contains` দিয়ে `exit` এর মতো command detect করা যায়।