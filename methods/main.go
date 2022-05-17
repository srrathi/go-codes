package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance no super no parent
	hitesh := User{"Hitesh", "test@gmail.com", true, 23}
	hitesh.GetStatus()
	hitesh.NewEmail()
	fmt.Println(hitesh)
	fmt.Printf("Details : %+v\n", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("Status of user is", user.Status)
}

func (user User) NewEmail() {
	user.Email = "test@google.com"
	fmt.Println("Email of the user is: ", user.Email)
}
