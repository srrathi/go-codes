package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"courseplatform"`
	Password string   `json:"-"`
	Tags     []string `json:"coursetags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON data")
	EncodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"Reactjs", 499, "web", "12345678", []string{"web-dev", "frontend", "javascript"}},
		{"Nodejs", 599, "app", "12345678", []string{"web-dev", "backend", "javascript"}},
		{"Golang", 499, "desktop", "12345678", nil},
	}

	finalJson, err := json.MarshalIndent(locCourses, "", "    ")
	// finalJson, err := json.Marshal(locCourses)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Data is %s\n", finalJson)
}
