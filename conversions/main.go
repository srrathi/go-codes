package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our World"
	fmt.Println(welcome)

	fmt.Printf("Please rate our pizza between 1 and 5 : ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error occurred", err)
		panic(err)
	}
	fmt.Println(numRating + 4)
}
