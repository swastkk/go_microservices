package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

// Here we have initialised a func init which will initate the opts variable.
func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./static/*.html",
	}

	rnd = renderer.New(opts)
}

// Defined the function named helloHandler which will render the page and the methods
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	rnd.HTML(w, http.StatusOK, "hello", nil)
}

// Defined formHandler function which will handle the form route
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v ", err)
		return
	}
	// fmt.Fprintf(w, "Post request succesfull! \n")
	// name := r.FormValue("name")
	// email := r.FormValue("email")
	rnd.HTML(w, http.StatusOK, "form", nil)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	// rnd = renderer.New(opts)
	fmt.Printf("Server starting on port 7000\n")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
