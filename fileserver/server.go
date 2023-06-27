package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	fmt.Printf("Server is up and running...")
	http.ListenAndServe(":5000", nil)
}
