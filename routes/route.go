package routes

import (
	"basictrade-gading/controllers"
	"basictrade-gading/middlewares"
	"basictrade-gading/repositories"
	"basictrade-gading/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteSession(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	adminRepo := repositories.NewAdminRepository(db)
	productRepo := repositories.NewProductRepository(db)

	adminService := services.NewAdminService(adminRepo)
	productService := services.NewProductService(productRepo)

	adminController := controllers.NewAdminController(adminService)
	productController := controllers.NewProductController(productService)

	adminRoute := router.Group("auth")
	{
		adminRoute.Use(middlewares.CORSMiddleware())
		adminRoute.POST("/register", middlewares.ValidationRequest("register-auth"), adminController.RegisterAdmin)
		adminRoute.POST("/login", middlewares.ValidationRequest("login-auth"), adminController.LoginAdmin)
	}

	productRoute := router.Group("products")
	{
		productRoute.Use(middlewares.CORSMiddleware(), middlewares.Authentication())
		productRoute.POST("/", middlewares.ValidationRequest("create-product"),productController.CreateProduct)
	}

	return router
}
