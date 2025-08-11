package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! this is my Go Web Server")
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server starting on: 8080... ")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
