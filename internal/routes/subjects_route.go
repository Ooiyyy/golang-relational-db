package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func SubjectsRoute(r *gin.Engine, subjectHandler *handler.SubjectsHandler) {
	classesRoute := r.Group("/api/v1/")
	{
		classesRoute.GET("/subjects", subjectHandler.GetAllSubjects)
		classesRoute.GET("/subject/:id", subjectHandler.GetSubjectByID)
		classesRoute.POST("/subject", subjectHandler.CreateSubjects)
		classesRoute.PUT("/subject/:id", subjectHandler.UpdateSubject)
		classesRoute.DELETE("/subject/:id", subjectHandler.DeleteSubject)
	}
}
