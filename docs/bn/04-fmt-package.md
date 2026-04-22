# পাঠ ০৪: Go-তে fmt প্যাকেজ

এই নোটে [examples/04-fmt-package/main.go](../../examples/04-fmt-package/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello Go Developers")

	name := "Chitano Kumar"
	age := 32
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	fmt.Println("My name is", name, "and I am ", age, "years old")
}
```

## এই পাঠে যা কভার করা হয়েছে

- `fmt` প্যাকেজ import ও ব্যবহার করা
- `fmt.Print`, `fmt.Printf`, এবং `fmt.Println`-এর পার্থক্য
- `%s`, `%d` এর মতো format verb ব্যবহার করে formatted string তৈরি
- স্বয়ংক্রিয় space এবং newline

## ব্যাখ্যা

1. `import "fmt"`  
   Go-এর built-in formatting ও I/O প্যাকেজ import করে। এটি console-এ print করতে এবং user input নিতে ব্যবহৃত হয়।

2. `fmt.Print("Hello Go Developers")`  
   string টি হুবহু print করে — শেষে কোনো newline যোগ হয় না এবং arguments-এর মধ্যে কোনো space যোগ হয় না।

3. `fmt.Printf("My name is %s and I am %d years old\n", name, age)`  
   Formatted string print করে। Format verb-গুলো variable দিয়ে প্রতিস্থাপিত হয়:
   - `%s` → string মান (`name`)
   - `%d` → integer মান (`age`)
   - `\n` → নতুন লাইন character

4. `fmt.Println("My name is", name, "and I am ", age, "years old")`  
   সব argument space দিয়ে আলাদা করে print করে এবং শেষে স্বয়ংক্রিয়ভাবে newline যোগ করে।

## প্রচলিত format verb-সমূহ

| Verb  | ধরন     | বিবরণ                              |
|-------|---------|------------------------------------|
| `%s`  | string  | সাধারণ string                      |
| `%d`  | integer | দশমিক integer                      |
| `%f`  | float   | দশমিক সংখ্যা                       |
| `%v`  | যেকোনো  | যেকোনো মানের default format        |
| `%+v` | struct  | field নামসহ struct                 |
| `%T`  | যেকোনো  | মানের type                         |
| `%t`  | bool    | Boolean (`true` অথবা `false`)      |
| `\n`  | —       | Newline (`Printf`-এর ভিতরে ব্যবহার) |

## উদাহরণ আউটপুট

```text
Hello Go DevelopersMy name is Chitano Kumar and I am 32 years old
My name is Chitano Kumar and I am  32 years old
```

> লক্ষ্য করো: `fmt.Print` newline যোগ করে না, তাই `Printf`-এর output একই লাইনে শুরু হয়।

## দ্রুত সারসংক্ষেপ

- Extra newline ছাড়া raw output-এর জন্য `fmt.Print` ব্যবহার করো।
- `%s`, `%d`-এর মতো verb দিয়ে formatted output-এর জন্য `fmt.Printf` ব্যবহার করো।
- সবচেয়ে সহজ output-এর জন্য `fmt.Println` ব্যবহার করো — এটি arguments-এর মধ্যে space এবং শেষে newline যোগ করে।
