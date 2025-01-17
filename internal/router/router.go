package router

import (
	"github.com/bigxxby/digital-travel-test/internal/api/middleware"
	authRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/auth"
	authService "github.com/bigxxby/digital-travel-test/internal/api/service/auth"
	authController "github.com/bigxxby/digital-travel-test/internal/api/transport/auth"

	userRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()

	// Initialize repositories, services, and controllers
	userRepo := userRepo.NewUserRepo(db)
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, userRepo)
	authController := authController.NewAuthController(authService)

	// Create groups and routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.GET("/whoami", middleware.AuthMiddleware(), authController.Whoami)
	}

	// Serve Swagger UI
	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	return router, nil
}
