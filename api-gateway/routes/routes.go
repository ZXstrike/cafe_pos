package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Route struct to hold the endpoint and its corresponding target service
type Route struct {
	Path   string
	Target string
}

// Initialize the routes
var routes = []Route{
	{"/users", "http://10.1.0.3"},  // User Service
	{"/menu", "http://10.1.0.4"},   // Menu Service
	{"/orders", "http://10.1.0.5"}, // Order Service
	{"/payment", "http://10.1.6."}, // Payment Service
}

// InitRoutes initializes the routes and proxies them to the target services
func InitRoutes() {
	for _, route := range routes {
		target, err := url.Parse(route.Target)
		if err != nil {
			log.Fatal("Error parsing URL: ", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		log.Printf("Proxying %s to %s\n", route.Path, route.Target)

		http.HandleFunc(route.Path+"/", func(w http.ResponseWriter, r *http.Request) {
			trimR := strings.Replace(r.URL.Path, route.Path, "", 1)

			r.URL.Path = trimR

			proxy.ServeHTTP(w, r)
		})
	}
}
