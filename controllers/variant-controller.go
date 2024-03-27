package controllers

import (
	"basictrade-gading/models"
	"basictrade-gading/services"
	"net/http"

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