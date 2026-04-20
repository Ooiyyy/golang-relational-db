package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func StudentRoute(r *gin.Engine, studentHandler *handler.StudentHandler, studentRelationalHandler *handler.StudentRelationalHandler) {
	// Grup rute yang membutuhkan otorisasi untuk mengakses Students
	studentsRoute := r.Group("/api/v1")
	{
		studentsRoute.GET("/students", studentHandler.GetAll)
		studentsRoute.GET("/students/:id", studentHandler.GetByID)

		studentsRoute.POST("/students", studentHandler.Create)
		studentsRoute.PUT("/students/:id", studentHandler.Update)
		studentsRoute.DELETE("/students/:id", studentHandler.Delete)

		// relasi
		studentsRoute.GET("/students/detail", studentRelationalHandler.GetAllDetails)
		studentsRoute.GET("/students/grades", studentRelationalHandler.GetStudentsGrade)
		studentsRoute.GET("/students/avg", studentRelationalHandler.GetStudentsAvg)

	}
}
