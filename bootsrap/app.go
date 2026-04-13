package bootstrap

import (
	"golang-sekolah/config"
	"golang-sekolah/internal/handler"
	"golang-sekolah/internal/repository/mysql"
	"golang-sekolah/internal/routes"
	"golang-sekolah/internal/service"
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
	todoRepo := mysql.NewTodoRepository(db)
	subjectRepo := mysql.NewSubjectRepo(db)

	userService := service.NewUserService(userRepo)
	todoService := service.NewTodoService(todoRepo)
	subjectService := service.NewSubjectService(subjectRepo)

	authHandler := handler.NewAuthHandler(userService, jwtSecret)
	todoHandler := handler.NewTodoHandlers(todoService)
	subjectHandler := handler.NewSubjectHandler(subjectService)

	r := routes.SetupRouter(
		authHandler,
		todoHandler,
		subjectHandler,
		jwtSecret,
	)

	// 4) Mulai HTTP server. Setelah ini request masuk via router Gin.
	r.Run(":" + port)
}
