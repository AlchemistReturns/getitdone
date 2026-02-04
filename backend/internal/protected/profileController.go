package protected

import (
	"example.com/getitdone/internal/models"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"name":  user.(models.User).Name,
		"email": user.(models.User).Email,
	})
}
