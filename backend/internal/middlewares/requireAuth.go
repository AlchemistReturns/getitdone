package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"example.com/getitdone/database"
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// 1. Get cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth cookie found"})
		return
	}

	// 2. Parse and Validate (Handles 'exp' and 'valid' automatically)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// 3. Centralized validation check
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// 4. Extract Claims and Find User
	claims, ok := token.Claims.(jwt.MapClaims)
	var user models.User

	if !ok || database.DB.First(&user, claims["sub"]).Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// 5. Success
	c.Set("user", user)
	c.Next()
}
