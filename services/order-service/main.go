package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/order-services/config"
	"github.com/zxstrike/cafe-pos/order-services/db"
	"github.com/zxstrike/cafe-pos/order-services/routes"

	"gorm.io/gorm"
)

func main() {

	// LoadConfig()
	config, er := config.LoadConfig()
	if er != nil {
		log.Fatalf("Error loading config: %v", er)
	}

	// InitializeDatabase()
	db, er := db.MySQLConnect(&config.MySQL)
	if er != nil {
		log.Fatalf("Error connecting to database: %v", er)
	}

	// StartServer()
	StartServer(config.ServePort, db)
}

// StartServer starts the server
func StartServer(Port string, db *gorm.DB) {
	router := gin.New()

	// Add middlewares
	router.Use(gin.Recovery(), gin.Logger())

	// Initialize app routes
	routes.InitRoutes(router)

	// Create an HTTP server using the Gin router
	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// <-ctx.Done()
	// fmt.Println("Timeout of 5 seconds.")

	fmt.Println("Server exited")
}
