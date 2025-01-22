package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zxstrike/cafe-pos/api-gateway/routes"
)

func main() {
	// Load routes
	routes.InitRoutes()

	// Start the server
	fmt.Println("API Gateway running on port 80...")
	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))
}
