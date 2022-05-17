package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models for data
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"courseAuthor"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db

var courses []Course

// middlewares
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to building Apis in Golang")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "123", CourseName: "Reactjs", CoursePrice: 499, Author: &Author{Fullname: "Sitaram", Website: "go.dev.com"}})
	courses = append(courses, Course{CourseId: "321", CourseName: "Nodejs", CoursePrice: 399, Author: &Author{Fullname: "Sitaram", Website: "go.dev.com"}})

	// routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", DeleteAllCourses).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":9001", r))
}

// controllers

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Golang</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	for _, oldCourse := range courses {
		if oldCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode(oldCourse)
			return
		}
	}

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Invalid Course id")
	return
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(params["id"] + " Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Invalid course id")
	return
}

func DeleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete All Courses")
	courses = nil
	json.NewEncoder(w).Encode("All courses are deleted")
	return
}
