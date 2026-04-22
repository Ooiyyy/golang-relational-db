package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func ClassRoute(r *gin.Engine, classHandler *handler.ClassHandler) {
	classesRoute := r.Group("/api/v1/")
	{
		classesRoute.GET("/classes", classHandler.GetAllClass)
		classesRoute.GET("/class/:id", classHandler.GetClassByID)
		classesRoute.POST("/class", classHandler.CreateClass)
		classesRoute.PUT("/class/:id", classHandler.UpdateClass)
		classesRoute.DELETE("/class/:id", classHandler.DeleteClass)
	}
}
