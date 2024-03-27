package controllers

import (
	"basictrade-gading/models"
	"basictrade-gading/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductController struct {
	productService services.IProductService
}

func NewProductController(productService services.IProductService) *ProductController {
	productController := ProductController{}
	productController.productService = productService
	return &productController
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {

	productCreateRequest := models.ProductRequest{}

	adminData := ctx.MustGet("adminData").(jwt.MapClaims)
	productCreateRequest.AdminId = uint(adminData["id"].(float64))

	if ctx.ContentType() == "multipart/form-data" {
		if err := ctx.Bind(&productCreateRequest); err != nil {
			log.Error().Msg(err.Error())
			models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		models.ResponseError(ctx, "Request should be form in multipart/form-data", http.StatusInternalServerError)
		return
	}

	resultData, err := c.productService.CreateProduct(&productCreateRequest)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	dataResponse := models.ProductResponse{
		UUID:     resultData.UUID,
		Name:     resultData.Name,
		ImageUrl: resultData.ImageUrl,
		AdminId:  resultData.AdminID,
	}

	models.ResponseSuccessWithData(ctx, dataResponse)

}

func (c *ProductController) GetProduct(ctx *gin.Context) {

	uuidProduct := ctx.Param("uuid")
	if uuidProduct == "" {
		models.ResponseError(ctx, "Product UUID cannot be empty", http.StatusBadRequest)
		return
	}

	resultData, err := c.productService.GetProduct(uuidProduct)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	dataResponse := models.ProductResponse{
		UUID:     resultData.UUID,
		Name:     resultData.Name,
		ImageUrl: resultData.ImageUrl,
		AdminId:  resultData.AdminID,
	}

	models.ResponseSuccessWithData(ctx, dataResponse)

}
