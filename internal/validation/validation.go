package validation

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func errorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fieldError.Field())
	case "number":
		return fmt.Sprintf("Field %s must be a number", fieldError.Field())
	case "startswith":
		return fmt.Sprintf("Field %s must start with %s", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("Field %s must greater than %s", fieldError.Field(), fieldError.Param())
	case "max":
		return fmt.Sprintf("Field %s must less than %s", fieldError.Field(), fieldError.Param())
	case "len":
		return fmt.Sprintf("Field %s must be %s characters", fieldError.Field(), fieldError.Param())
	case "email":
		return fmt.Sprintf("Field %s must be a valid email", fieldError.Field())
	case "maxFileSize":
		return fmt.Sprintf("Field %s must be an image", fieldError.Field())
	case "customDateFormat":
		return fmt.Sprintf("Field %s must be a valid date format, it should be YYYY-MM-DD", fieldError.Field())
	case "oneof":
		return fmt.Sprintf("Field %s must be one of the following: %s", fieldError.Field(), fieldError.Param())
	}

	return fieldError.Error()
}

func ValidateRequest(str interface{}) interface{} {
	validate := validator.New()
	validate.RegisterValidation("customDateFormat", customDateFormat)

	if err := validate.Struct(str); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errors := make([]string, len(ve))
			i := 0
			for _, fieldError := range ve {
				errors[i] = errorMessage(fieldError)
				i++
			}
			return errors
		}
		// Handle unexpected error type
		return err.Error()
	}
	return nil
}

func customDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, dateStr)
	if !match {
		return false
	}

	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
