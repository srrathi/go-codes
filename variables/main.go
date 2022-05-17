package main

import "fmt"

const LoginTokken string = "dcfagvybhuijxkdjbhbv acbhxjkln" // public variable

func main() {
	var userName string = "rohit rathi"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n\n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n\n", isLoggedIn)

	var friends int = 800
	fmt.Println(friends)
	fmt.Printf("Variable is of type: %T \n\n", friends)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n\n", smallInt)

	var smallFloat float32 = 23355.3456789765435678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n\n", smallFloat)

	var largeFloat float64 = 23355.3456789765435678
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n\n", largeFloat)

	var friendsList int
	fmt.Println(friendsList)
	fmt.Printf("Variable is of type: %T \n\n", friendsList)

	friends2 := 56
	fmt.Println(friends2)
	fmt.Printf("Variable is of type: %T \n\n", friends2)

	fmt.Println(LoginTokken)
	fmt.Printf("Variable is of type: %T \n\n", LoginTokken)

}
