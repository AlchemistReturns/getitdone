package public

import (
	"github.com/gin-gonic/gin"
)

// Home is a simple public handler that returns a welcome message.
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home",
	})
}
