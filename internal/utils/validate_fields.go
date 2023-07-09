package utils

import (
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct[T any](payload T) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
func ShowErrors(errors []*models.ErrorResponse) models.ErrorDetail {
	var stringError string
	count := 0
	for _, err := range errors {
		count += 1
		stringError += "Field: " + err.Field + " [tag-error: " + err.Tag + "]"
		if count < len(errors) {
			stringError += ", "
		}
	}
	return models.ErrorDetail{
		TypeError:        "Error fields",
		ErrorDescription: stringError,
	}
}
