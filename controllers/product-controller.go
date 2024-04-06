package controllers

import (
	"basictrade-gading/models"
	"basictrade-gading/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
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

func (c *ProductController) GetProducts(ctx *gin.Context) {

	offsetParam, ok := ctx.GetQuery("offset")
	if !ok {
		models.ResponseError(ctx, "Offset cannot be empty", http.StatusBadRequest)
		return
	}

	limitParam, ok := ctx.GetQuery("limit")
	if !ok {
		models.ResponseError(ctx, "Limit cannot be empty", http.StatusBadRequest)
		return
	}

	nameFilter, _ := ctx.GetQuery("name")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		log.Error().Msg("Invalid offset value")
		models.ResponseError(ctx, "Invalid offset value", http.StatusInternalServerError)
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 0 {
		log.Error().Msg("Invalid offset value")
		models.ResponseError(ctx, "Invalid limit value", http.StatusInternalServerError)
		return
	}

	result, err := c.productService.GetProducts(limit, offset, nameFilter)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccessWithData(ctx, result)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {

	productUpdateRequest := models.ProductUpdateRequest{}

	uuid := ctx.Param("uuid")

	if ctx.ContentType() == "multipart/form-data" {
		if err := ctx.Bind(&productUpdateRequest); err != nil {
			log.Error().Msg(err.Error())
			models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		models.ResponseError(ctx, "Request should be form in multipart/form-data", http.StatusInternalServerError)
		return
	}

	resultData, err := c.productService.UpdateProduct(uuid, &productUpdateRequest)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	dataResponse := models.ProductResponse{
		UUID: resultData.UUID,
		Name: resultData.Name,
		ImageUrl: resultData.ImageUrl,
		AdminId: resultData.AdminID,
	}

	models.ResponseSuccessWithData(ctx, dataResponse)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context){

	uuid := ctx.Param("uuid")

	err := c.productService.DeleteProduct(uuid)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccess(ctx, "Delete product is success")
}
