package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int

	fmt.Println(ptr)
	fmt.Printf("Type of ptr is %T\n", ptr)

	num := 2342

	var ptr2 = &num
	fmt.Println(*ptr2)
	fmt.Println(ptr2)
	fmt.Printf("Type of ptr is %T\n", ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println(*ptr2)
}
