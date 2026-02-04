package main

import (
	"example.com/getitdone/database"
	"example.com/getitdone/internal/auth"
	"example.com/getitdone/internal/middlewares"
	"example.com/getitdone/internal/protected"
	"example.com/getitdone/internal/public"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect()
	database.Migrate()
}

func main() {

	router := gin.Default()
	api := router.Group("/api/v1")
	private := api.Group("/protected")

	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)

	api.GET("/", public.Home)

	private.Use(middlewares.RequireAuth)
	{
		private.GET("/profile", protected.Profile)
	}

	router.Run(":8080")
}
