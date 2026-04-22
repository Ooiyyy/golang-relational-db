package bootstrap

import (
	"golang-relational-db/config"
	"golang-relational-db/internal/handler"
	"golang-relational-db/internal/repository/mysql"
	"golang-relational-db/internal/routes"
	"golang-relational-db/internal/service"
)

func Run() {
	// 1) Load konfigurasi runtime (port, JWT secret, kredensial DB) dari env.
	cfg := config.LoadEnv()
	port := cfg.App.AppPort
	jwtSecret := cfg.App.JWTSecret
	// 2) Buka koneksi DB pool sekali saat startup, dipakai bersama di semua request.
	db := config.ConnectDB(cfg.DB)

	// 3) Rangkai dependensi dari lapisan paling bawah ke atas:
	// repository -> service -> handler.
	userRepo := mysql.NewUserRepository(db)
	subjectRepo := mysql.NewSubjectRepo(db)
	scoreRepo := mysql.NewScoreRepo(db)
	studentRepo := mysql.NewStudentRepo(db)
	studentRelationalRepo := mysql.NewStudentRelationalRepo(db)
	classRepo := mysql.NewClassRepo(db)
	subjectsRepo := mysql.NewSubjectsRepo(db)
	gradeRepo := mysql.NewGradesRepo(db)
	teacherRepo := mysql.NewTeacherRepo(db)
	activityLogRepo := mysql.NewActivityLogRepo(db)

	userService := service.NewUserService(userRepo)
	subjectService := service.NewSubjectService(subjectRepo)
	scoreService := service.NewScoreService(scoreRepo)
	studentService := service.NewStudentService(studentRepo, classRepo)
	studentRelationalService := service.NewStudentRelationalService(studentRelationalRepo)
	classServ := service.NewClassService(classRepo, teacherRepo)
	subjectsService := service.NewSubjectsService(subjectsRepo, teacherRepo)
	gradeServ := service.NewGradesService(gradeRepo, studentRepo, subjectsRepo)
	teacherServ := service.NewTeacherService(teacherRepo)

	authHandler := handler.NewAuthHandler(userService, jwtSecret)
	subjectHandler := handler.NewSubjectHandler(subjectService)
	scoreHandler := handler.NewScoreHandler(scoreService)
	studentHandler := handler.NewStudentHandler(studentService, activityLogRepo)
	studentRelationalHandler := handler.NewStudentRelationalHandler(studentRelationalService, activityLogRepo)
	classHandler := handler.NewClassHandler(classServ, activityLogRepo)
	subjectsHandler := handler.NewSubjectsHandler(subjectsService, activityLogRepo)
	gradeHandler := handler.NewGradesHandler(*gradeServ, activityLogRepo)
	teacherHandler := handler.NewTeacherHandler(teacherServ, activityLogRepo)

	r := routes.SetupRouter(
		authHandler,
		subjectHandler,
		scoreHandler,
		studentHandler,
		studentRelationalHandler,
		classHandler,
		subjectsHandler,
		gradeHandler,
		teacherHandler,
		jwtSecret,
	)

	// 4) Mulai HTTP server. Setelah ini request masuk via router Gin.
	r.Run(":" + port)
}
