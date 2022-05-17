package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance no super no parent
	hitesh := User{"Hitesh", "test@gmail.com", true, 23}
	fmt.Println(hitesh)
	fmt.Printf("Details : %+v\n", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

