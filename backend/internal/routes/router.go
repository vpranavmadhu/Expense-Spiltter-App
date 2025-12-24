package routes

import (
	"esapp/internal/handlers"
	"esapp/internal/middleware"
	"esapp/internal/repository"
	"esapp/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	return r
}
