package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my About Page!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status": "API is running"}`)

	/*fullName := "John Oche Abalike"
	parts := strings.Split(fullName, " ")

	if len(parts) == 3 {
		firstName := parts[0]
		middleName := parts[1]
		lastName := parts[2]

		fmt.Println("First Name", firstName)
		fmt.Println("Middle Name", middleName)
		fmt.Println("Last Name", lastName)
	} else {
		fmt.Println("Name format is not First Middle Last")
	}*/
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "User Profile Page for the user ID: %s", userID)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/api", apiHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
