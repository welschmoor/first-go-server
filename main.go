package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>Welcome blah</h2>")
}

func handleNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>List of Notes:</h2><ul>  	<li>Note number 1</li>	<li>Note number 2</li><li>Note number 3</li>	</ul>")
}

func handleBlogs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>Blogs</h2>")
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404 page")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/notes", handleNotes)
	http.HandleFunc("/blogs", handleBlogs)
	http.HandleFunc("*", handleBlogs)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
