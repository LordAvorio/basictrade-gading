package controllers

import (
	"basictrade-gading/models"
	"basictrade-gading/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService services.IAdminService
}

func NewAdminController(adminService services.IAdminService) *AdminController {
	adminController := AdminController{}
	adminController.adminService = adminService
	return &adminController
}

func (c *AdminController) RegisterAdmin(ctx *gin.Context) {

	adminRegisterRequest := models.AdminCreateRequest{}

	if ctx.ContentType() == "multipart/form-data" {
		if err := ctx.Bind(&adminRegisterRequest); err != nil {
			models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		models.ResponseError(ctx, "Request should be form in multipart/form-data", http.StatusInternalServerError)
		return
	}

	checkValidation := adminRegisterRequest.ValidationRegister()
	if len(checkValidation) > 0 {
		models.ResponseErrorWithData(ctx, "Validation Error", http.StatusBadRequest, checkValidation)
		return
	}

	adminRegisterData := models.Admin{
		Name:     adminRegisterRequest.Name,
		Email:    adminRegisterRequest.Email,
		Password: adminRegisterRequest.Password,
	}

	err := c.adminService.RegisterAdmin(&adminRegisterData)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccess(ctx, "Admin register is successfull")

}
