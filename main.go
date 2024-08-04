package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Data represents the data we want to return
type Data struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

func main() {
	// Define the data we want to return
	data := Data{
		Message: "Hello, World!",
		Value:   42,
	}

	// Define the HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}

	// Create the HTTP server and define the endpoint
	http.HandleFunc("/data", handler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}