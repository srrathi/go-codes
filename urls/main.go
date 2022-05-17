package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/todos/1?course=reactjs&paymentId=vbnbnbnhbnhgb"

func main() {
	fmt.Println("Welcome to handling urls")
	fmt.Println(myUrl)
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	query := result.Query()
	fmt.Printf("type of query params is %T\n", query)
	fmt.Println(query.Get("course"))
	fmt.Printf("type of reactjs params without square bracket is %T\n", query.Get("course"))
	fmt.Println(query["course"][0])
	fmt.Printf("type of reactjs params in square bracket is %T\n", query["course"])
	fmt.Println(query.Get("paymentId"))

	for _, val := range query {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "aws.com",
		Path: "/srrathi/profile",
		RawQuery: "user=admin&page=4",
	}

	fmt.Printf("Type of url maker: %T\n\n", partsOfUrl)

	fullUrl := partsOfUrl.String()
	fmt.Println(fullUrl)
}
