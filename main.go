package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Listening on port 8080\n")
	if er := http.ListenAndServe(":8080", nil); er != nil {
		log.Fatal(er)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 404ing any not allow routes
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Handing any non GET requests
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Form Name: %s Address: %s", name, address)
}
