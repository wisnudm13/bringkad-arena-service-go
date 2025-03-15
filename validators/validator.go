package validators

import (
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		minLen   = false
		hasUpper = false
		hasLower = false
		hasNum   = false
		hasSpec  = false
	)

	if len(password) >= 8 {
		minLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNum = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpec = true
		}
	}

	return minLen && hasLower && hasUpper && hasNum && hasSpec

}

func RegisterRequestCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", PasswordValidator)
	}
}
