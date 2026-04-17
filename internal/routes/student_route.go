package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func StudentRoute(r *gin.Engine, studentHandler *handler.StudentHandler) {
	// Grup rute yang membutuhkan otorisasi untuk mengakses Students
	studentsRoute := r.Group("/api/v1")
	{
		// Akses Publik (Siswa & Guru)
		studentsRoute.GET("/students", studentHandler.GetAll)
		studentsRoute.GET("/students/:id", studentHandler.GetByID)

		// Akses Terbatas (Hanya Guru)
		studentsRoute.POST("/students", studentHandler.Create)
		studentsRoute.PUT("/students/:id", studentHandler.Update)
		studentsRoute.DELETE("/students/:id", studentHandler.Delete)
	}
}
