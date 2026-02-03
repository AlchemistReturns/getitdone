package auth

import (
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	c.JSON(200, gin.H{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}
