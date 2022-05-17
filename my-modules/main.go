package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello from mod in golang")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func Greeter() {
	fmt.Println("Welcome to our website")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my Golang served site</h1>"))
}
