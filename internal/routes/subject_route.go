package routes

import (
	"golang-sekolah/internal/handler"
	"golang-sekolah/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SubjectRoute(r *gin.Engine, subjectHandler *handler.SubjectHandler, jwtSecret string) {
	// Grup rute yang membutuhkan otorisasi untuk mengakses Subject
	subjectRoutes := r.Group("/subject")
	subjectRoutes.Use(middleware.JWT(jwtSecret))
	{
		// Akses Publik (Siswa & Guru)
		subjectRoutes.GET("/", middleware.RoleCheck("guru", "siswa"), subjectHandler.GetAll)
		subjectRoutes.GET("/:id", middleware.RoleCheck("guru", "siswa"), subjectHandler.GetByID)

		// Akses Terbatas (Hanya Guru)
		subjectRoutes.POST("/", middleware.RoleCheck("guru"), subjectHandler.Create)
		subjectRoutes.PATCH("/:id", middleware.RoleCheck("guru"), subjectHandler.Update)
		subjectRoutes.DELETE("/:id", middleware.RoleCheck("guru"), subjectHandler.Delete)
	}
}
