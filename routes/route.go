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

	adminService := services.NewAdminService(adminRepo)

	adminController := controllers.NewAdminController(adminService)

	adminRoute := router.Group("auth")
	{
		adminRoute.POST("/register", middlewares.CORSMiddleware() ,adminController.RegisterAdmin)
	}

	return router
}