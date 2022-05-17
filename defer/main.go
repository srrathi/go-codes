package main

import "fmt"

func main() {
	// LIFO ORDER THAT MEANS LAST IN FIRST OUT
	defer fmt.Println("Hello from Defer")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("World from Defer")
	fmt.Println("World from Defer")
	fmt.Println("World from Defer")
	fmt.Println("World from Defer")
	fmt.Println("World from Defer")
	fmt.Println("World from Defer")
	myDefer()

	fmt.Println("World from Defer")
	fmt.Println("World from Defer")

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
