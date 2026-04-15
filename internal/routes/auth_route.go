package routes

import (
	"golang-relational-db/internal/handler"
	"golang-relational-db/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine, authHandler *handler.AuthHandler, jwtSecret string) {

	// Public endpoint: request bisa masuk tanpa token.
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	// Protected endpoint: request harus lolos JWT middleware dulu.
	protected := r.Group("/")
	protected.Use(middleware.JWT(jwtSecret))
	{
		protected.GET("/profile", authHandler.Profile)
		protected.POST("/logout", authHandler.Logout)
	}
}
