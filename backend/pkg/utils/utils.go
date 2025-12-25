package utils

import "github.com/gin-gonic/gin"

func NotImplemented(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "not implemented yet",
	})
}
