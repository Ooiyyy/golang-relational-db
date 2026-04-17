package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler, subjectHandler *handler.SubjectHandler, scoreHandler *handler.ScoreHandler, studentHandler *handler.StudentHandler, jwtSecret string) *gin.Engine {
	// gin.Default() menambahkan logger + recovery middleware bawaan.
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// Daftarkan endpoint per modul.
	AuthRoute(r, authHandler, jwtSecret)
	SubjectRoute(r, subjectHandler, jwtSecret)
	ScoreRoute(r, scoreHandler, jwtSecret)
	StudentRoute(r, studentHandler)

	// Router ini dipanggil oleh bootstrap untuk menerima request dari client.
	return r
}
