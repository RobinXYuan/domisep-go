package auth

import (
	"strings"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// const passwordFormatPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@_-?*&]).{8,}$"

//PasswordValidator 密码校验
func PasswordValidator(field validator.FieldLevel) bool {
	specialChars := "@!_-=:*"
	letters := 0
	hasNumber, hasUpperCase, hasSpecialChar := false, false, false

	for _, c := range field.Field().String() {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
			letters++
		case unicode.IsUpper(c):
			hasUpperCase = true
			letters++
		case strings.Contains(specialChars, string(c)):
			hasSpecialChar = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false
		}
	}

	if letters >= 8 && hasNumber && hasUpperCase && hasSpecialChar {
		return true
	}

	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", PasswordValidator)
	}
}
