package middleware

import (
	"RealTimeChatApplication/utils"

	"github.com/gin-gonic/gin"
)

func MidwareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claim, err := utils.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user", claim)
		c.Next()
	}
}
