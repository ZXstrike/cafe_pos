package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zxstrike/cafe-pos/api-gateway/routes"
)

func logRequest(r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL)
}

func main() {
	// Load routes
	routes.InitRoutes()

	// Start the server
	fmt.Println("API Gateway running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
