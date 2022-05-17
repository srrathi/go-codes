package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	fmt.Printf("%v\n", greeter)

	result := proAdded(3, 50, 49)
	value, msg := doubleReturn(3, 5, 6, 7)
	// result := addition(3, 49)
	fmt.Println(value, msg)
	fmt.Println(result)
	greeter()
}

func greeter() {
	fmt.Println("Namastey from Golang")
}

func addition(a int, b int) int {
	return a + b
}

func proAdded(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func doubleReturn(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "String also gets returned"
}
