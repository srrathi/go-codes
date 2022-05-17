package main

import "fmt"

func main() {
	fmt.Println("Conditional Statements in Golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println(("Less than 10"))
	} else {
		fmt.Println("Greater than 10")
	}

}
