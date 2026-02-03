package auth

import (
	"example.com/getitdone/database"
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	//Bind the request body to the user struct
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user.Password = string(hashedPassword)

	//Save the user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to save user",
		})
		return
	}

	//Return the user
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}
