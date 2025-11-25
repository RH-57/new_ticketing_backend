package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	// Handle validator.v10
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()

			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)

			case "email":
				errorsMap[field] = "Invalid email format"

			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())

			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())

			case "oneof":
				errorsMap[field] = fmt.Sprintf("%s must be one of: %s", field, fieldError.Param())

			default:
				errorsMap[field] = "Invalid value"
			}
		}
	}

	// Handle Duplicate Entry GORM
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists"
			}
			if strings.Contains(err.Error(), "code") {
				errorsMap["Code"] = "Code already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorsMap["Error"] = "Record not found"
		}
	}

	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
