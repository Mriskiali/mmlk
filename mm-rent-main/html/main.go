// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the root page!")
}

func handleHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../mm-rent-main 1/mm-rent-main/html/index.html")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/html", handleHTML)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
