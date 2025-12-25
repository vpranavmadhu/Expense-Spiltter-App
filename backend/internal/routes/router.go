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

	//repos
	userRepo := repository.NewUserRepository(db)
	groupReop := repository.NewGroupRepository(db)

	//services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	groupService := services.NewGroupService(groupReop)

	//handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	groupHandler := handlers.NewGroupHandler(groupService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	api := r.Group("/api")
	{
		api.Use(middleware.AuthMiddleware())
		api.GET("/me", userHandler.GetMe)
		api.POST("/creategroup", groupHandler.CreateGroup)
	}
	return r
}
