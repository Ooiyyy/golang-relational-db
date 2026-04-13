// digunakan untuk menyeragamkan response (tapi belum dipake dulu)

package utils

import (
	"github.com/gin-gonic/gin"
)

// SendErrorResponse membantu menyeragamkan format respon error ke pengguna
func SendErrorResponse(c *gin.Context, statusCode int, errorTitle string, errorMessage interface{}) {
	c.JSON(statusCode, gin.H{
		"error":   errorTitle,
		"message": errorMessage, // errorMessage bisa berupa string (err.Error()) atau array (untuk validation error)
	})
}

// SendSuccessResponse (Bonus) membantu menyeragamkan format respon sukses dengan data
func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	if data != nil {
		c.JSON(statusCode, gin.H{
			"message": message,
			"data":    data,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": message,
	})
}
