package main

import (
	"example.com/getitdone/database"
	"example.com/getitdone/internal/auth"
	"example.com/getitdone/internal/middlewares"
	"example.com/getitdone/internal/protected"
	"example.com/getitdone/internal/public"
	"github.com/gin-gonic/gin"
)

// init runs before main. It's used here to set up the database connection
// and run any pending migrations to ensure the DB schema is up to date.
func init() {
	database.Connect()
	database.Migrate()
}

// main is the entry point of the application.
func main() {

	// Initialize the Gin router (default comes with Logger and Recovery middleware)
	router := gin.Default()

	// Create a route group for API version 1.
	// All routes in this group will start with /api/v1
	api := router.Group("/api/v1")

	// Define public routes (no authentication required)
	// POST /register - Creates a new user
	api.POST("/register", auth.Register)
	// POST /login - Authenticates a user and returns a JWT
	api.POST("/login", auth.Login)
	// GET / - Simple home route
	api.GET("/", public.Home)

	// Define a private route group.
	// All routes here will be protected by the RequireAuth middleware.
	private := api.Group("/protected")

	// Use the RequireAuth middleware for this group.
	// This ensures that any request to /api/v1/protected/* must have a valid JWT.
	private.Use(middlewares.RequireAuth)
	{
		// GET /profile - Returns the logged-in user's profile information
		private.GET("/profile", protected.Profile)
	}

	// Start the server on port 8080
	router.Run(":8080")
}
