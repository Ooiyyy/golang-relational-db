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

	userService := service.NewUserService(userRepo)
	subjectService := service.NewSubjectService(subjectRepo)
	scoreService := service.NewScoreService(scoreRepo)

	authHandler := handler.NewAuthHandler(userService, jwtSecret)
	subjectHandler := handler.NewSubjectHandler(subjectService)
	scoreHandler := handler.NewScoreHandler(scoreService)

	r := routes.SetupRouter(
		authHandler,
		subjectHandler,
		scoreHandler,
		jwtSecret,
	)

	// 4) Mulai HTTP server. Setelah ini request masuk via router Gin.
	r.Run(":" + port)
}
