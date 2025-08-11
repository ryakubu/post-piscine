package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// --- Structs ---
type product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

// --- Handlers ---
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var products product

	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real app, save to a database here

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product created successfully",
		"product": products,
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our mini API!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	response := map[string]string{"id": userID, "name": "John Deo"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	response := map[string]string{"category": category, "message": "Product list"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	response := map[string]string{"query": query, "page": page, "results": "Some search results"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Page not found")
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// --- Main ---
func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler).Methods("GET")
	r.HandleFunc("/products/{category}", productsHandler).Methods("GET")
	r.HandleFunc("/search", searchHandler).Methods("GET")
	r.HandleFunc("/products", createProductHandler).Methods("POST")

	// 404 Handler
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Middleware
	r.Use(loggingMiddleWare)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
