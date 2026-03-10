package httphelper

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse ensures all data is wrapped in a "data" key
func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, gin.H{
		"status": "success",
		"data":   payload,
		"error":  nil,
	})
}

// ErrorResponse ensures all errors look the same
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status": "error",
		"data":   nil,
		"error":  message,
	})
}
