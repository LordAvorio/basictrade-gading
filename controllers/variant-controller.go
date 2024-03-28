package controllers

import (
	"basictrade-gading/models"
	"basictrade-gading/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type VariantController struct {
	variantService services.IVariantService
}

func NewVariantController(variantService services.IVariantService) *VariantController {
	variantController := VariantController{}
	variantController.variantService = variantService
	return &variantController
}

func(c *VariantController)CreateVariant(ctx *gin.Context){

	variantCreateRequest := models.VariantRequest{}

	if ctx.ContentType() == "multipart/form-data" {
		if err := ctx.Bind(&variantCreateRequest); err != nil {
			log.Error().Msg(err.Error())
			models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		models.ResponseError(ctx, "Request should be form in multipart/form-data", http.StatusInternalServerError)
		return
	}

	resultData, err := c.variantService.CreateVariant(&variantCreateRequest)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	dataResponse := models.VariantResponse{
		UUID: resultData.UUID,
		VariantName: resultData.VariantName,
		Quantity: resultData.Quantity,
		ProductID: resultData.ProductID,
	}

	models.ResponseSuccessWithData(ctx, dataResponse)

}

func (c *VariantController) GetVariant(ctx *gin.Context) {

	uuidVariant := ctx.Param("uuid")
	if uuidVariant == "" {
		models.ResponseError(ctx, "Variant UUID cannot be empty", http.StatusBadRequest)
		return
	}

	resultData, err := c.variantService.GetVariant(uuidVariant)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	dataResponse := models.VariantResponse{
		UUID: resultData.UUID,
		VariantName: resultData.VariantName,
		Quantity: resultData.Quantity,
		ProductID: resultData.ProductID,
	}

	models.ResponseSuccessWithData(ctx, dataResponse)

}

func (c *VariantController) GetVariants(ctx *gin.Context) {

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

	nameFilter, _ := ctx.GetQuery("variantName")

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

	result, err := c.variantService.GetVariants(limit, offset, nameFilter)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return		
	}

	models.ResponseSuccessWithData(ctx, result)
}