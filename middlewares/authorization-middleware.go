package middlewares

import (
	"basictrade-gading/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AdminAuthorization(category string, db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")

		tokenData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(tokenData["id"].(float64))

		switch category {
		case "product":
			productModel := models.Product{}
			err := db.Select("admin_id").Where("uuid = ?", uuid).First(&productModel).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Admin not found",
					"message": err.Error(),
				})
				return
			}

			if productModel.AdminID != adminID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this API",
				})
				return
			}
		case "variant":
			variantModel := models.Variant{}
			productModel := models.Product{}

			err := db.Select("product_id").Where("uuid = ?", uuid).First(&variantModel).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Product not found",
					"message": err.Error(),
				})
				return
			}

			err = db.Select("admin_id").Where("id = ?", variantModel.ProductID).First(&productModel).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Admin not found",
					"message": err.Error(),
				})
				return
			}

			if productModel.AdminID != adminID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this API",
				})
				return
			}

		}

		ctx.Next()
	}

}
