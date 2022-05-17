package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello from files")
	content := "This will go in file"

	file, err := os.Create("./file.txt")
	defer file.Close()
	checkNilError(err)

	length, err := io.WriteString(file, content)
	fmt.Printf("File address is %v\n", file)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	readFile("./file.txt")
	// file.Close()
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data in file is :", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
