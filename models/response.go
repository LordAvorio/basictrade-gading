package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseModelWithData struct {
	Data interface{} `json:"data"`
	ResponseModel
}

func ResponseSuccessWithData(ctx *gin.Context, data interface{}) {

	response := ResponseModelWithData{}
	response.Code = http.StatusOK
	response.Data = data
	response.Message = "Success"

	ctx.JSON(http.StatusOK, response)

}

func ResponseSuccessCreated(ctx *gin.Context, data interface{}) {

	response := ResponseModelWithData{}
	response.Code = http.StatusCreated
	response.Data = data
	response.Message = "Created"

	ctx.JSON(http.StatusCreated, response)

}

func ResponseSuccess(ctx *gin.Context, message string) {

	response := ResponseModel{}
	response.Code = http.StatusOK
	response.Message = message

	ctx.JSON(http.StatusOK, response)

}

func ResponseError(ctx *gin.Context, err string, code int) {

	response := ResponseModel{}
	response.Code = code
	response.Message = err

	ctx.JSON(code, response)

}

func ResponseErrorWithData(ctx *gin.Context, err string, code int, data any){
	response := ResponseModelWithData{}
	response.Code = code
	response.Message = err
	response.Data = data

	ctx.JSON(code, response)
}
