package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func GradeRoute(r *gin.Engine, gradeHandler *handler.GradesHandler,
	relationalHandler *handler.StudentRelationalHandler) {
	gradesRoute := r.Group("/api/v1/")
	{
		gradesRoute.GET("/grades", relationalHandler.GetStudentsGrade)
		gradesRoute.POST("/grade", gradeHandler.CreateGrades)
	}
}
