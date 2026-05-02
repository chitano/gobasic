# Go Learning Journey

This repository documents my Go learning journey.

I use it to:
- practice concepts with small examples
- write and test beginner-to-intermediate Go code
- track what I learn over time
- share progress publicly

## Why this repo exists

I am learning Go by building consistent habits:
- read and learn a concept
- write code for that concept
- revise and improve code quality
- keep notes for future reference

## Repository structure

Current structure:

```text
.
|-- README.md
|-- docs/
|   |-- 01-hello-go-developers.md
|   |-- 02-variable.md
|   |-- 03-constand.md
|   |-- 04-fmt-package.md
|   |-- 07-custome-data-type.md
|   |-- cli-app-fmt-pkg.md
|   |-- cli-app-todolist.md
|   |-- bn/
|       |-- 01-hello-go-developers.md
|       |-- 02-variable.md
|       |-- 03-constand.md
|       |-- 04-fmt-package.md
|       |-- 07-custome-data-type.md
|       |-- cli-app-fmt-pkg.md
|       |-- cli-app-todolist.md
|-- examples/
|   |-- 01-hello-go-developers/
|   |   |-- main.go
|   |-- 02-variable/
|   |   |-- main.go
|   |-- 03-constant/
|   |   |-- main.go
|   |-- 04-fmt-package/
|   |   |-- main.go
|   |-- 07-custome-data-type/
|       |-- main.go
|-- projects/
|   |-- cli-app/
|   |   |-- fmt-pkg/
|   |   |   |-- go.mod
|   |   |   |-- main.go
|   |   |-- todolist/
|   |   |   |-- main.go
|   |   |   |-- tasks.json
|   |-- grpc/
|   |-- rest-api/
```

## Docs index

Learning notes available right now:

- [01 - Hello Go Developers](docs/01-hello-go-developers.md)
- [02 - Variable](docs/02-variable.md)
- [03 - Constant](docs/03-constand.md)
- [04 - The fmt Package](docs/04-fmt-package.md)
- [07 - Custom Data Types](docs/07-custome-data-type.md)
- [Project - CLI App with fmt Package](docs/cli-app-fmt-pkg.md)
- [Project - Todo List CLI](docs/cli-app-todolist.md)

## Docs index (Bangla)

- [০১ - Hello Go Developers](docs/bn/01-hello-go-developers.md)
- [০২ - Variable](docs/bn/02-variable.md)
- [০৩ - Constant](docs/bn/03-constand.md)
- [০৪ - fmt প্যাকেজ](docs/bn/04-fmt-package.md)
- [০৭ - Custom Data Type](docs/bn/07-custome-data-type.md)
- [প্রজেক্ট - fmt প্যাকেজ দিয়ে CLI অ্যাপ](docs/bn/cli-app-fmt-pkg.md)
- [প্রজেক্ট - Todo List CLI](docs/bn/cli-app-todolist.md)

## Projects

Run commands for the current CLI projects:

### fmt-pkg

```bash
cd projects/cli-app/fmt-pkg
go run main.go
```

### todolist

```bash
cd projects/cli-app/todolist
go run main.go list
go run main.go add "Learn Go"
go run main.go done 1
go run main.go delete 1
```

## Learning goals

- Build a strong foundation in Go syntax and core features
- Understand pointers, structs, interfaces, and error handling
- Practice writing clean and readable Go code
- Learn goroutines and channels with practical examples
- Build small projects and gradually increase complexity

## How to run code

From this folder:

```bash
go run main.go
```

If I add subfolders later, each folder may have its own runnable file or package.

## Progress log

I will keep updating this section.

- [ ] Setup complete
- [ ] Basic syntax practice
- [ ] Functions and control flow
- [ ] Arrays, slices, and maps
- [ ] Structs and methods
- [ ] Interfaces and error handling
- [ ] Goroutines and channels
- [ ] Mini project 1

## Notes

This repo is intentionally beginner-friendly and evolving. Code quality and structure will improve as I learn more.

## Connect

If you are learning Go too, feel free to explore, suggest improvements, or share ideas.

