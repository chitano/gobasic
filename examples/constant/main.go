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
