package utils

import (
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Custom validation error messages
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "password":
		return "Password must be at least 8 characters, include uppercase, lowercase, and a special character"
	default:
		return "Invalid input"
	}
}

// validate request body TODO: add form data validation
func RequestValidator(c *gin.Context, rq interface{}) error {
	if err := c.ShouldBindJSON(rq); err != nil {
		if errors.Is(err, io.EOF) { //empty request body or invalid format
			BadRequestMessage(c, "Invalid Request Body")
			return err
		}

		// custom message validation
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var errMsg string

			for _, fe := range ve {
				errMsg = getErrorMessage(fe)
				break
			}
			UnprocessableEntityMessage(c, errMsg)
			return err
		}

		BadRequestMessage(c, err.Error())
		return err

	}

	return nil
}
