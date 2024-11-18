package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/zxstrike/cafe-pos/api-gateway/middlewares"
)

// Route struct to hold the endpoint and its corresponding target service
type Route struct {
	Path   string
	Target string
}

// Initialize the routes
var routes = []Route{
	{"/users", "http://localhost:8081"},   // User Service
	{"/orders", "http://localhost:8082"},  // Order Service
	{"/menu", "http://localhost:8083"},    // Menu Service
	{"/payment", "http://localhost:8084"}, // Payment Service
}

// InitRoutes initializes the routes and proxies them to the target services
func InitRoutes() {
	for _, route := range routes {
		target, err := url.Parse(route.Target)
		if err != nil {
			log.Fatal("Error parsing URL: ", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		// Apply AuthMiddleware to routes
		http.Handle(route.Path+"/", middlewares.AuthMiddleware(http.StripPrefix(route.Path, proxy)))
	}
}
