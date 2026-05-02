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
