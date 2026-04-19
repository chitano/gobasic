package main

import "fmt"

func main() {
	var name string = "Chitano Kumar" //Explicitly define type.
	fmt.Println("My name is", name)

	age := 34 //Auto detected type or type inferance in Golang

	fmt.Println("My age is", age)

	email := "chitanokumar@gmail.com"

	count, _ := fmt.Println(email)

	fmt.Println("Number of character in my email", count)

}
