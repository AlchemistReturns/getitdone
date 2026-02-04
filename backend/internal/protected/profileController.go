package protected

import (
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
)

// Profile returns the authenticated user's profile information.
func Profile(c *gin.Context) {

	// Get the user object that was set in the RequireAuth middleware.
	// We ignore the exists boolean (_) because the middleware guarantees it exists.
	user, _ := c.Get("user")

	// Type assertion: Convert the interface{} type back to models.User struct
	// to access its fields (Name, Email).
	c.JSON(200, gin.H{
		"name":  user.(models.User).Name,
		"email": user.(models.User).Email,
	})
}
