package auth

import (
	"os"
	"time"

	"example.com/getitdone/database"
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	//Get Email and password from request body
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Look up the user
	var user models.User
	database.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}

	//Check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid password",
		})
		return
	}

	//Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	//Sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	//Send back the token as a cookie
	c.SetCookie("Authorization", tokenString, 60*60*24, "/", "localhost", false, true)
	c.JSON(200, gin.H{})
}
