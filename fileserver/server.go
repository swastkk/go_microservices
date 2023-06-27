package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	fmt.Printf("FileServer is up and running at PORT 5000")
	http.ListenAndServe(":5000", nil)
}

// Here is the detailed implementation of the Code that is done above!

// func main() {
// 	dir := http.Dir("./public/")
// 	fmt.Printf("This is the dir: %s\n", dir)
// 	dirHandler := http.FileServer(dir)
// 	http.Handle("/", dirHandler)
// 	fmt.Print("Fileserver is up and running at port 5000...\n ")
// 	http.ListenAndServe(":5000", nil)

// }
