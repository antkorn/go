package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web!!!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Listen to port 3000...")
	http.ListenAndServe(":3000", nil)

}
