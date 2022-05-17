package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to our World"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating : ")
	
	// comma ok // err ok
	input, err := reader.ReadString('\n')
	fmt.Printf("Thanks for rating %s\n", input)
	fmt.Println(err)
	
}