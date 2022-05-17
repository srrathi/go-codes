package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go Server")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl string = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	fmt.Printf("Response Status Code %v\n", response.StatusCode)
	fmt.Printf("Response Headers %v\n", response.Header)
	// fmt.Println(response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	byteCount, _ := responseString.Write(dataBytes)
	fmt.Printf("Response is %v\n", string(dataBytes))
	fmt.Printf("Response bytecount is %v\n", byteCount)
	fmt.Println(responseString.String())
}
