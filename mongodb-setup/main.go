package main

import (
	"fmt"
	"log"
	"mongodb-setup/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Mongodb API with Golang")
	r := router.Router()
	fmt.Println("Server is getting started")
	fmt.Println("Listening at port 9001")
	fmt.Println("http://localhost:9001/")
	log.Fatal(http.ListenAndServe(":9001", r))
}
