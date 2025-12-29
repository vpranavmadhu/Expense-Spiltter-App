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
	expenseRepo := repository.NewExpenseRepository(db)
	settlementRepo := repository.NewSettlementRepository(db)

	//services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	groupService := services.NewGroupService(groupReop, userRepo)
	expenseService := services.NewExpenseService(expenseRepo, groupReop, settlementRepo)

	//handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	groupHandler := handlers.NewGroupHandler(groupService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

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
		api.GET("/groups/:groupId", groupHandler.GetGroupByID)
		api.POST("/groups/:groupId/addmember", groupHandler.AddMember)
		api.GET("/groups", groupHandler.ListGroups)
		api.GET("/groups/:groupId/members", groupHandler.ListMembers)
		api.POST("/createexpense", expenseHandler.CreateExpense)
		api.GET("/groups/:groupId/expenses", expenseHandler.ListExpensesWithShare)
		api.GET("/groups/:groupId/balances", expenseHandler.CalculateBalances)
		api.POST("/settlements", expenseHandler.MarkAsPaid)
	}
	return r
}
