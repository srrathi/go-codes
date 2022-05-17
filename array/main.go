package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "peach"
	fruitList[2] = "mango"
	fruitList[3] = "banana"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is :", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is :", vegList)

}
