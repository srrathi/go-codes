package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Grapes", "Banana")
	fmt.Println("fruit list is", fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println("fruit list is", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 232
	highScores[1] = 989
	highScores[2] = 121
	highScores[3] = 543

	highScores = append(highScores, 666, 909)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{}

	courses = append(courses, "reactjs", "nodejs", "mongodb")
	fmt.Println(courses)

	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
