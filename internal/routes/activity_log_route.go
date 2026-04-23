package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func LogRoute(r *gin.Engine, logHandler *handler.ActivityLogHandler) {
	logRoute := r.Group("/api/v1/")
	{
		logRoute.GET("/logs", logHandler.GetAllLog)
		logRoute.GET("/log/:id", logHandler.GetLogByID)
		// logRoute.POST("/class", logHandler.CreateClass)
		// logRoute.PUT("/class/:id", logHandler.UpdateClass)
		// logRoute.DELETE("/class/:id", logHandler.DeleteClass)
	}
}
