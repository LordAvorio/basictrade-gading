package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitGoValidation() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateImageHeader(fileImageHeader *multipart.FileHeader, fieldName string) map[string]string {

	resultError := map[string]string{}

	maxSize := int64(1 << 20)
	if fileImageHeader.Size > maxSize {
		resultError[fieldName] = "File size exceeds maximum allowed size (1 MB)"
	}

	formatImage := []string{".jpg", ".jpeg", ".png", ".webp"}
	extensionImage := strings.ToLower(filepath.Ext(fileImageHeader.Filename))
	validExt := false
	for _, ext := range formatImage {
		if ext == extensionImage {
			validExt = true
			break
		}
	}

	if !validExt {
		resultError[fieldName] = "Extension must be on .jpg, .jpeg, .png, .webp"
	}

	return resultError
}

func ValidationMessage(field string, tag string) string {
	message := ""

	switch tag {
	case "required":
		message = fmt.Sprintf("%s is required", field)
	case "email":
		message = "Wrong format email"
	case "lte":
		message = fmt.Sprintf("%s value is too long", field)
	case "gte":
		message = fmt.Sprintf("%s value is too short", field)
	case "image":
		message = "Wrong format image"
	}

	return message
}
