package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

// AuthMiddleware to check JWT token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Look for Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := authHeader[len("Bearer "):]
		// Here you would validate the token (JWT or similar)
		// For simplicity, we'll skip the actual validation for now

		fmt.Println("Authenticated with token:", token)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
