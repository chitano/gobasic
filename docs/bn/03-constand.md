# পাঠ ০৩: Go-তে Constants

এই নোটে [examples/constant/main.go](../../examples/constant/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

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

## এই পাঠে যা কভার করা হয়েছে

- `const` দিয়ে constant declare করা
- const block-এ একাধিক constant group করা
- constant এবং variable এর পার্থক্য
- `:=` দিয়ে variable shadowing

## ব্যাখ্যা

1. `const (...)`
   এখানে package-level তিনটি constant declare করা হয়েছে: `Hostname`, `DBHost`, `AppName`।

2. `Hostname` বড় হাতের অক্ষর দিয়ে শুরু
   Go-তে এটি exported নামে ধরা হয় (অন্য package থেকে access করা যায়), যদি non-`main` package-এ ব্যবহার করা হয়।

3. `fmt.Println(Hostname)`
   constant value `http://localhost` প্রিন্ট করে।

4. `//Hostname = "Value not replaceable"`
   এই লাইন uncomment করলে compile error হবে।
   কারণ constant reassign করা যায় না।

5. `Hostname := "mydomain.com"`
   `main` ফাংশনের ভিতরে একই নামে নতুন local variable তৈরি হয়।
   এটি constant `Hostname`-কে এই scope-এ shadow (hide) করে।

6. `Hostname = "Now it replaceable"`
   এখন এটি কাজ করছে, কারণ এখানে `Hostname` বলতে local variable বোঝানো হচ্ছে, constant নয়।

7. `fmt.Print(Hostname)`
   update করা local variable value প্রিন্ট করে।

## উদাহরণ আউটপুট

```text
http://localhost
Now it replaceable
```

## গুরুত্বপূর্ণ নোট

- constant declare করার পর value পরিবর্তন করা যায় না।
- shadowing সম্ভব, কিন্তু constant আর local variable-এ একই নাম ব্যবহার করলে readability কমতে পারে।
- `DBHost` এবং `AppName` declare করা হলেও `main`-এ এখনো ব্যবহার করা হয়নি।

## দ্রুত সারসংক্ষেপ

- স্থির value-এর জন্য `const` ব্যবহার করো
- local variable declare করতে `:=` ব্যবহার করো
- একই নামে local variable package-level constant-কে shadow করতে পারে