package models

import (
	"basictrade-gading/utils"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"time"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null; type:varchar(155)"`
	Name      string `gorm:"not null; type:varchar(100); unique"`
	ImageUrl  string `gorm:"not null; type:varchar(255)"`
	AdminID   uint
	Admin     *Admin
	Variants  []Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type ProductRequest struct {
	Name    string               `form:"name" validate:"required,lte=100"`
	Image   multipart.FileHeader `form:"file" validate:"required"`
	AdminId uint
}

type ProductUpdateRequest struct {
	Name    string               `form:"name" validate:"required,lte=100"`
	Image   multipart.FileHeader `form:"file"`
}

type ProductResponse struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	AdminId  uint   `json:"admin_id"`
}

func (pr *ProductRequest) ValidationProductCreate() map[string]string {

	errorMessage := map[string]string{}
	err := utils.Validate.Struct(pr)

	errorValidationImage := utils.ValidateImageHeader(&pr.Image, "file")
	if len(errorValidationImage) > 0 {
		for key, value := range errorValidationImage {
			errorMessage[key] = value
		}
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = utils.ValidationMessage(err.Field(), err.Tag())
		}
	}

	return errorMessage
}

func (pur *ProductUpdateRequest) ValidationProductUpdate() map[string]string {

	errorMessage := map[string]string{}
	err := utils.Validate.Struct(pur)

	if pur.Image.Filename != "" {
		errorValidationImage := utils.ValidateImageHeader(&pur.Image, "file")
		if len(errorValidationImage) > 0 {
			for key, value := range errorValidationImage {
				errorMessage[key] = value
			}
		}
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = utils.ValidationMessage(err.Field(), err.Tag())
		}
	}

	return errorMessage

}
