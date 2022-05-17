package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to go POST request")
	PerformPostRequest()

}

func PerformPostRequest() {
	const myurl string = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
	{
		"title": "foo", "body": "bar", "userId": 1
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	fmt.Printf("Response Status Code %v\n", response.StatusCode)
	fmt.Printf("Response Headers %v\n", response.Header.Get("Content-Length"))
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
