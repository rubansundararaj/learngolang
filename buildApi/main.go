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

//Model for course
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

//middleware

func (c *Course) IsEmpty() bool {
	//return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Autocodegen is running from Golang")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseID: "2", CourseName: "React Js", CoursePrice: 100, Author: &Author{FullName: "Ruban", Website: "autocodegen.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourses).Methods("GET")
	r.HandleFunc("/create", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serves home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Legendary auto code generator</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through find matching id and return the response.Body

	for _, c := range courses {
		if c.CourseID == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here we create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	//handle {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}

	// generate unique id, string, and course
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("course")
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here we update one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	// grab params from req
	params := mux.Vars(r)

	//loop through slice, find object with id, replace the id

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// Send reponse when id is not found
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	// grab params from req
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			return
		}
	}

}
