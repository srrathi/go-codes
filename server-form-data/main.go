package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to sending form data in go")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl string = "https://jsonplaceholder.typicode.com/posts"

	data := url.Values{}

	data.Add("title", "foo")
	data.Add("body", "bar")
	data.Add("userId", "201")

	response, err := http.PostForm(myurl, data)
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
