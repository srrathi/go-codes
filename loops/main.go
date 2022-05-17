package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday", "Monday"}
	days = append(days, "Valentines")
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("%v) %v\n", index+1, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		fmt.Println("Value is :", rogueValue)
		rogueValue++
	}
}
