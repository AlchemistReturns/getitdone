package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RequireAuth is a middleware that validates the JWT token in cookies.
// It ensures that only authenticated users can access the protected routes.
func RequireAuth(c *gin.Context) {
	// 1. Get the Authorization cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No auth cookie found"})
		return
	}

	// 2. Parse and Validate the token
	// jwt.Parse will check the signature and expiration ('exp') automatically.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (matches the one used in Login)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to sign the token
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// 3. Centralized validation check
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// 4. Extract Claims and Create User context
	// Convert claims to MapClaims to access custom fields (sub, name, email)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Extract data safely
	sub, ok := claims["sub"].(float64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	name, ok := claims["name"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user name in token"})
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user email in token"})
		return
	}

	user := models.User{
		Name:  name,
		Email: email,
	}
	user.ID = uint(sub)

	// 5. Success
	c.Set("user", user)
	c.Next()
}
