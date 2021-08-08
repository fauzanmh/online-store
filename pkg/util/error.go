package util

import (
	"fmt"

	"github.com/fauzanmh/online-store/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func errorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return helper.PqError(err)
	}
	return helper.CommonError(err)
}

// * isPqError used to check error if error is pg error
func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

func switchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for idx, err := range castedObject {
			field := ToSnakeCase(err.Field())

			// Change Field Name
			switch field {
			}

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "email":
				message = fmt.Sprintf("%s must be valid email address",
					field)
			case "gt":
				message = fmt.Sprintf("%s value must be greater than %s",
					field, err.Param())
			default:
				message = err.Error()
			}

			if idx == 0 {
				break
			}

		}
	}
	return
}
