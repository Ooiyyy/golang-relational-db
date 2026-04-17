package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type List struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
	Errors  interface{} `json:"errors"`
}

type Meta struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

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

// success response
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

// error response
func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
}

// lisr response
func ListResponse(message string, data interface{}, meta Meta) List {
	return List{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
		Errors:  nil,
	}
}
