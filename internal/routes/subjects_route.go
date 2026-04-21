package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func SubjectsRoute(r *gin.Engine, subjectHandler *handler.SubjectsHandler) {
	classesRoute := r.Group("/api/v1/")
	{
		classesRoute.GET("/subjects", subjectHandler.GetAllSubjects)
		classesRoute.POST("/subject", subjectHandler.CreateSubjects)
	}
}
