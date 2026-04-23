package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Errors      interface{} `json:"errors"`
	DateTime    string      `json:"datetime"`
	DateTimeWIB string      `json:"datetime_wib"`
	Timezone    string      `json:"timezone"`
}

type List struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Meta        Meta        `json:"meta"`
	Errors      interface{} `json:"errors"`
	DateTime    string      `json:"datetime"`
	DateTimeWIB string      `json:"datetime_wib"`
	Timezone    string      `json:"timezone"`
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
	nowUtc := time.Now().UTC()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowWib := nowUtc.In(loc)
	return Response{
		Success:     true,
		Message:     message,
		Data:        data,
		Errors:      nil,
		DateTime:    nowUtc.Format(time.RFC3339),
		DateTimeWIB: nowWib.Format("2006-01-02 15:04:05"),
		Timezone:    "UTC",
	}
}

// error response
func ErrorResponse(message string, errors interface{}) Response {
	nowUtc := time.Now().UTC()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowWib := nowUtc.In(loc)
	return Response{
		Success:     false,
		Message:     message,
		Data:        nil,
		Errors:      errors,
		DateTime:    nowUtc.Format(time.RFC3339),
		DateTimeWIB: nowWib.Format("2006-01-02 15:04:05"),
		Timezone:    "UTC",
	}
}

// lisr response
func ListResponse(message string, data interface{}, meta Meta) List {
	nowUtc := time.Now().UTC()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowWib := nowUtc.In(loc)
	return List{
		Success:     true,
		Message:     message,
		Data:        data,
		Meta:        meta,
		Errors:      nil,
		DateTime:    nowUtc.Format(time.RFC3339),
		DateTimeWIB: nowWib.Format("2006-01-02 15:04:05"),
		Timezone:    "UTC",
	}
}
