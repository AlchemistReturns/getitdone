package main

import (
	"fmt"

	"example.com/getitdone/database"
	"example.com/getitdone/internal/auth"
	"example.com/getitdone/internal/public"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect()
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)

	api.GET("/", public.Home)

	router.Run(":8080")
}
