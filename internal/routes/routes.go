package routes

import (
	"golang-relational-db/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler,
	subjectHandler *handler.SubjectHandler,
	scoreHandler *handler.ScoreHandler,
	studentHandler *handler.StudentHandler,
	studentRelationalHandler *handler.StudentRelationalHandler,
	classHandler *handler.ClassHandler,
	subjectsHandler *handler.SubjectsHandler,
	gradeHandler *handler.GradesHandler,
	teacherHandler *handler.TeacherHandler,
	logHandler *handler.ActivityLogHandler,
	jwtSecret string) *gin.Engine {
	// gin.Default() menambahkan logger + recovery middleware bawaan.
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// Daftarkan endpoint per modul.
	AuthRoute(r, authHandler, jwtSecret)
	SubjectRoute(r, subjectHandler, jwtSecret)
	ScoreRoute(r, scoreHandler, jwtSecret)
	StudentRoute(r, studentHandler, studentRelationalHandler)
	ClassRoute(r, classHandler)
	SubjectsRoute(r, subjectsHandler)
	GradeRoute(r, gradeHandler, studentRelationalHandler)
	TeacherRoute(r, teacherHandler)
	LogRoute(r, logHandler)

	// Router ini dipanggil oleh bootstrap untuk menerima request dari client.
	return r
}
