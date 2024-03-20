package models

import (
	"basictrade-gading/utils"
	"fmt"
	"time"
	"github.com/go-playground/validator/v10"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null; type:varchar(155)"`
	Name      string `gorm:"not null; type:varchar(100)"`
	Email     string `gorm:"not null; type:varchar(155); unique"`
	Password  string `gorm:"not null; type:varchar(155);"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type AdminCreateRequest struct {
	Name     string `form:"name" validate:"required,lte=100"`
	Email    string `form:"email" validate:"required,lte=155,email"`
	Password string `form:"password" validate:"required,lte=155"`
}

type AdminResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func validationMessage(field string, tag string) string {
	message := ""

	switch tag {
	case "required":
		message = fmt.Sprintf("%s is required", field)
	case "email":
		message = "Wrong format email"
	case "lte":
		message = fmt.Sprintf("%s value is too long", field)
	}

	return message
}

func (acr *AdminCreateRequest) ValidationRegister() map[string]string {

	errorMessage := map[string]string{}

	err := utils.Validate.Struct(acr)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = validationMessage(err.Field(), err.Tag())
		}

	}

	return errorMessage
}
