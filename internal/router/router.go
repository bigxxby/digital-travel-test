package router

import (
	"github.com/bigxxby/digital-travel-test/internal/api/middleware"
	authRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/auth"
	authService "github.com/bigxxby/digital-travel-test/internal/api/service/auth"
	authController "github.com/bigxxby/digital-travel-test/internal/api/transport/auth"
	"github.com/go-redis/redis"

	productRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/product"
	userRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/user"
	productService "github.com/bigxxby/digital-travel-test/internal/api/service/product"
	productController "github.com/bigxxby/digital-travel-test/internal/api/transport/product"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, redisDB *redis.Client) (*gin.Engine, error) {
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

	productRepo := productRepo.NewProductRepo(db)
	productService := productService.NewProductService(productRepo, userRepo)
	productController := productController.NewProductController(productService)

	product := router.Group("/product")
	{
		product.POST("/", middleware.AuthMiddleware(), productController.CreateProduct)
		product.PUT("/:productId", middleware.AuthMiddleware(), productController.UpdateProduct)
		product.DELETE("/:productId", middleware.AuthMiddleware(), productController.DeleteProduct)
		product.GET("/:productId", middleware.AuthMiddleware(), productController.GetProductById)
		product.GET("/", middleware.AuthMiddleware(), productController.GetAllProducts)
	}

	// Serve Swagger UI
	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	return router, nil
}
