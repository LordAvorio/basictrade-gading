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
	variantRepo := repositories.NewVariantRepository(db)

	adminService := services.NewAdminService(adminRepo)
	productService := services.NewProductService(productRepo)
	variantService := services.NewVariantService(variantRepo, productRepo)

	adminController := controllers.NewAdminController(adminService)
	productController := controllers.NewProductController(productService)
	variantController := controllers.NewVariantController(variantService)

	adminRoute := router.Group("auth")
	{
		adminRoute.Use(middlewares.CORSMiddleware())
		adminRoute.POST("/register", middlewares.ValidationRequest("register-auth"), adminController.RegisterAdmin)
		adminRoute.POST("/login", middlewares.ValidationRequest("login-auth"), adminController.LoginAdmin)
	}

	productRoute := router.Group("products")
	{
		productRoute.Use(middlewares.CORSMiddleware())
		productRoute.POST("/", middlewares.Authentication(), middlewares.ValidationRequest("create-product"), productController.CreateProduct)
		productRoute.GET("/:uuid", productController.GetProduct)
		productRoute.GET("/", productController.GetProducts)
		productRoute.PUT("/:uuid", middlewares.Authentication(), middlewares.AdminAuthorization("product", db), middlewares.ValidationRequest("update-product"), productController.UpdateProduct)
		productRoute.DELETE("/:uuid", middlewares.Authentication(), middlewares.AdminAuthorization("product", db), productController.DeleteProduct)
	}

	variantRoute := router.Group("variants")
	{
		variantRoute.Use(middlewares.CORSMiddleware())
		variantRoute.POST("/", middlewares.Authentication(), middlewares.ValidationRequest("create-variant"), variantController.CreateVariant)
		variantRoute.GET("/:uuid", variantController.GetVariant)
		variantRoute.GET("/", variantController.GetVariants)
		variantRoute.PUT("/:uuid", middlewares.Authentication(), middlewares.AdminAuthorization("variant", db), middlewares.ValidationRequest("update-variant"), variantController.UpdateVariant)
	}

	return router
}
