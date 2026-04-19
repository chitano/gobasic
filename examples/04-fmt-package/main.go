package main

import "fmt"

func main() {
	// In most programming language we have built in packages for showing output to screen, console or any other channel. And olso a way to take input from users.
	// In go we have "fmt" package for standard input/output.
	//Simple output, you can just print/output anything.
	fmt.Print("Hello Go Developers")

	//formatted output, you can give a structure in your output.
	name := "Chitano Kumar"
	age := 32
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	//If you want a new line after your output then you can use it. Also it will add  space between all argruments.
	fmt.Println("My name is", name, "and I am ", age, "years old")

}
