# পাঠ ০৭: Go-তে Custom Data Type

এই নোটে [examples/07-custome-data-type/main.go](../../examples/07-custome-data-type/main.go) ফাইলের কোড ব্যাখ্যা করা হয়েছে।

## সোর্স কোড

```go
package main

import (
	"fmt"
	"regexp"
	"time"
)

type userID    uint32
type Email     string
type AddressId uint32
type Gender    string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Address struct {
	Id           AddressId
	AddressLine1 string
	// ...
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
	// 5 জন demo user তৈরি করা হয়েছে
	for _, u := range users {
		fmt.Printf("User: %-15s | Email: %-30s | Valid: %v\n", u.Name, u.Email, u.Email.IsValid())
	}
}
```

## এই পাঠে যা কভার করা হয়েছে

- primitive-এর উপর ভিত্তি করে custom named type তৈরি করা (`uint32`, `string`)
- custom `Gender` type দিয়ে typed constant
- একাধিক field সহ struct define করা (`Address`, `User`)
- custom type-এ method attach করা (method receiver)
- `regexp` দিয়ে email validation
- struct-এর slice
- `fmt.Printf` দিয়ে formatted output

## ব্যাখ্যা

### Custom type

```go
type userID    uint32
type Email     string
type AddressId uint32
type Gender    string
```

এগুলো বিদ্যমান primitive-এর উপর ভিত্তি করে নতুন ও আলাদা type। যদিও `Email` আসলে `string`-এর উপর তৈরি, explicit conversion ছাড়া `string` থেকে `Email`-এ assign করা যায় না। এটি type safety নিশ্চিত করে।

### Typed constant

```go
const (
	Male   Gender = "male"
	Female Gender = "female"
)
```

সর্বত্র `"male"` বা `"female"` raw string ব্যবহারের বদলে `Gender` type-এর named constant define করা হয়েছে। এটি কোড পড়তে সহজ করে এবং typo প্রতিরোধ করে।

### Struct

```go
type Address struct { ... }
type User struct { ... }
```

Struct সংশ্লিষ্ট field-গুলোকে একটি নামে একত্রিত করে। `User` struct-এ `Address` সরাসরি field হিসেবে embed করা হয়েছে (composition)।

### `Gender`-এর উপর method — `IsValid()`

```go
func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	}
	return false
}
```

Method হল এমন একটি function যার একটি receiver আছে। এখানে `g` হল `Gender` type-এর value receiver। Method টি শুধুমাত্র তখনই `true` return করে যখন gender টি defined constant-গুলোর মধ্যে একটি।

### Email validation — `IsValid()`

```go
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (e Email) IsValid() bool {
	return emailRegex.MatchString(string(e))
}
```

- `regexp.MustCompile` package level-এ একবার regex compile করে (দক্ষ — প্রতিবার call করার সময় পুনরায় compile হয় না)।
- Method টি match করার আগে `Email` type-কে plain `string`-এ convert করে।
- Regex টি `local@domain.tld` pattern enforce করে।

### Struct-এর slice

```go
users := []User{ ... }
```

5টি pre-populated demo record সহ `User` struct-এর একটি slice (dynamic list) তৈরি করে।

### Iterate ও print করা

```go
for _, u := range users {
	fmt.Printf("User: %-15s | Email: %-30s | Valid: %v\n", u.Name, u.Email, u.Email.IsValid())
}
```

- `range` slice-এর প্রতিটি element-এ iterate করে।
- `%-15s` string টি 15-character-wide column-এ left-align করে।
- `%v` যেকোনো মান তার default format-এ print করে (`true`/`false` bool-এর জন্য)।

## উদাহরণ আউটপুট

```text
User: Alice Johnson    | Email: alice.johnsonexample.com    | Valid: false
User: Bob Smith        | Email: bob.smith@example.com       | Valid: true
User: Carol Lee        | Email: carol.lee@example.com       | Valid: true
User: David Brown      | Email: david.brown@example.com     | Valid: true
User: Eva Green        | Email: eva.green@example.com       | Valid: true
```

> Alice-এর email-এ `@` symbol নেই, তাই `IsValid()` সঠিকভাবে `false` return করে।

## গুরুত্বপূর্ণ নোট

- Primitive-এর উপর ভিত্তি করে তৈরি custom type explicit conversion ছাড়া **interchangeable নয়**: `string(myEmail)`।
- `regexp.MustCompile` regex invalid হলে startup-এ panic করে — শুধুমাত্র compile-time-এ জানা pattern-এর জন্য ব্যবহার করো।
- Go-তে method receiver value receiver `(t Type)` বা pointer receiver `(t *Type)` হতে পারে। মান পরিবর্তন করতে হলে pointer receiver ব্যবহার করো।

## দ্রুত সারসংক্ষেপ

- আরও expressive ও type-safe কোডের জন্য `type` দিয়ে named type তৈরি করো।
- Method দিয়ে type-এ behaviour যুক্ত করো।
- Pattern-based validation-এর জন্য `regexp` ব্যবহার করো।
- Go-তে Struct + method হল object-oriented-style design-এর ভিত্তি।
