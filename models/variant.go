package models

import (
	"basictrade-gading/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey"`
	UUID        string `gorm:"not null; type:varchar(155)"`
	VariantName string `gorm:"not null; type:varchar(100);"`
	Quantity    int    `gorm:"not null; type:bigint"`
	ProductID   uint
	Product     *Product
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type VariantRequest struct {
	VariantName string `form:"variant_name" validate:"required",lte=100`
	Quantity    int    `form:"quantity" validate:"required",gte=0`
	UUID        string `form:"product_id" validate:"required"`
}

type VariantResponse struct {
	UUID        string `json:"uuid"`
	VariantName string `json:"variant_name"`
	Quantity    int    `json:"quantity"`
	ProductID   uint   `json:"product_id"`
}

func (pr *VariantRequest) ValidationVariantCreate() map[string]string {

	errorMessage := map[string]string{}
	err := utils.Validate.Struct(pr)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = utils.ValidationMessage(err.Field(), err.Tag())
		}
	}

	return errorMessage
}