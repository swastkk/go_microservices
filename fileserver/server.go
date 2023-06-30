package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// Here is the detailed implementation of the FileServer Code.

func main() {
	dir := http.Dir("./public/")
	fmt.Printf("This is the directory which is being served: %s\n", dir)
	dirHandler := http.FileServer(dir)
	http.Handle("/", dirHandler)
	fmt.Print("Fileserver is up and running at port 5000...\n ")
	u := url.PathEscape("http://localhost:5000")
	fmt.Printf("Check here: %v\n", u)
	http.ListenAndServe(":5000", nil)
}

// Can implement the code in the shorter format like this! Both works same.
// func main() {
// 	http.Handle("/", http.FileServer(http.Dir("./public/")))
// 	fmt.Printf("FileServer is up and running at PORT 5000")
// 	http.ListenAndServe(":5000", nil)
// }

// It took around 03:38:71(Min:Sec:MS) for sharing 1.21GB mp4 format file.
