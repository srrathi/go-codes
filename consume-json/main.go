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
	fmt.Println("Welcome to consuming JSON data")
	DecodeJson()
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`
	{
        "coursename": "Nodejs",
        "courseprice": 599,
        "courseplatform": "app",
        "coursetags": [
            "web-dev",
            "backend",
            "javascript"
        ]
    }
	`)

	var locCourse course

	checkValid := json.Valid(JsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(JsonDataFromWeb, &locCourse)
		fmt.Printf("%#v\n", locCourse)
		name := locCourse.Name
		price := locCourse.Price
		platform := locCourse.Platform
		tags := locCourse.Tags

		fmt.Println("Data stored in variables", name, price, platform, tags)
		fmt.Printf("%v\n", locCourse.Name)
		fmt.Printf("%v\n", locCourse)
		fmt.Printf("%T\n", locCourse)
	} else {
		fmt.Println("JSON was not valid")
	}
}
