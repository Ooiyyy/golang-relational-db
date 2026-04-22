package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func TeacherRoute(r *gin.Engine, teacherHandler *handler.TeacherHandler) {
	teacherRoute := r.Group("/api/v1/")
	{
		teacherRoute.GET("/teachers", teacherHandler.GetAll)
		teacherRoute.GET("/teacher/:id", teacherHandler.GetByID)
		teacherRoute.POST("/teacher", teacherHandler.Create)
		teacherRoute.PUT("/teacher/:id", teacherHandler.Update)
		teacherRoute.DELETE("/teacher/:id", teacherHandler.Delete)
	}
}
