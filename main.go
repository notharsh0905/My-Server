package main

import (
	"fmt"
	"net/http"

	"my_server/config"
	"my_server/handlers"
	"my_server/middleware"
)

func main() {
	// Initialize database connection setup
	config.InitDB()
	defer config.DB.Close()

	// Static route registration
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	// API Endpoint registrations wrapped by Middleware
	http.Handle("/create", middleware.LoggerMiddleware(http.HandlerFunc(handlers.CreateUserData)))
	http.Handle("/users", middleware.LoggerMiddleware(http.HandlerFunc(handlers.GetUserHandler)))

	fmt.Println("Server starting on port 8090 with professional structure...")
	http.ListenAndServe(":8090", nil)
}
