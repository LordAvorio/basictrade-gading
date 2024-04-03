package middlewares

import (
	"basictrade-gading/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidationRequest(section string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		errorMessage := map[string]string{}

		switch section {
		case "register-auth":
			adminRegisterRequest := models.AdminCreateRequest{}
			if err := ctx.Bind(&adminRegisterRequest); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			checkRegisterValidation := adminRegisterRequest.ValidationRegister()
			if len(checkRegisterValidation) > 0 {
				errorMessage = checkRegisterValidation
			}
		case "login-auth":
			adminLoginRequest := models.AdminLoginRequest{}
			if err := ctx.Bind(&adminLoginRequest); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			checkLoginValidation := adminLoginRequest.ValidationLogin()
			if len(checkLoginValidation) > 0 {
				errorMessage = checkLoginValidation
			}
		case "create-product":
			productCreateRequest := models.ProductRequest{}
			if err := ctx.Bind(&productCreateRequest); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			createProductValidation := productCreateRequest.ValidationProductCreate()
			if len(createProductValidation) > 0 {
				errorMessage = createProductValidation
			}
		case "create-variant":
			variantCreateRequest := models.VariantRequest{}
			if err := ctx.Bind(&variantCreateRequest); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			createVariantValidation := variantCreateRequest.ValidationVariantCreate()
			if len(createVariantValidation) > 0 {
				errorMessage = createVariantValidation
			}
		case "update-product":
			productUpdateRequest := models.ProductUpdateRequest{}
			if err := ctx.Bind(&productUpdateRequest); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			createProducttValidation := productUpdateRequest.ValidationProductUpdate()
			if len(createProducttValidation) > 0 {
				errorMessage = createProducttValidation
			}
		}

		if len(errorMessage) > 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorMessage)
			return
		}

		ctx.Next()
	}

}
