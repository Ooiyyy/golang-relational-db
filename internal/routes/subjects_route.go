package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func SubjectsRoute(r *gin.Engine, subjectHandler *handler.SubjectsHandler) {
	subjectsRoute := r.Group("/api/v1/")
	{
		subjectsRoute.GET("/subjects", subjectHandler.GetAllSubjects)
		subjectsRoute.GET("/subject/:id", subjectHandler.GetSubjectByID)
		subjectsRoute.POST("/subject", subjectHandler.CreateSubjects)
		subjectsRoute.PUT("/subject/:id", subjectHandler.UpdateSubject)
		subjectsRoute.DELETE("/subject/:id", subjectHandler.DeleteSubject)
	}
}
