package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["C#"] = "C Sharp"
	languages["CPP"] = "C Plus Plus"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")

	for key, value := range languages {
		fmt.Printf("Key is : %s and value is : %s\n", key, value)
	}
	fmt.Println(languages)
}
