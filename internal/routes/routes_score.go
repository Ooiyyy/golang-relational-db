package routes

import (
	"golang-sekolah/internal/handler"
	"golang-sekolah/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ScoreRoute(r *gin.Engine, scoreHandler *handler.ScoreHandler, jwtSecret string) {
	// Grup rute yang membutuhkan otorisasi untuk mengakses Subject
	scoreRoutes := r.Group("/score")
	scoreRoutes.Use(middleware.JWT(jwtSecret))
	{
		// Akses Publik (Siswa & Guru)
		scoreRoutes.GET("/", middleware.RoleCheck("guru", "siswa"), scoreHandler.GetAll)
		scoreRoutes.GET("/:id", middleware.RoleCheck("guru", "siswa"), scoreHandler.GetByID)

		// Akses Terbatas (Hanya Guru)
		scoreRoutes.POST("/", middleware.RoleCheck("guru"), scoreHandler.Create)
		scoreRoutes.PATCH("/:id", middleware.RoleCheck("guru"), scoreHandler.Update)
		scoreRoutes.DELETE("/:id", middleware.RoleCheck("guru"), scoreHandler.Delete)
	}
}
