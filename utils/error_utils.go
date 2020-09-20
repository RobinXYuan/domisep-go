package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

//FormatValidationError 参数校验错误
type FormatValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

const defaultReason = "Not valid value for field %s!"
const defaultRequiredReason = "Field %s is required!"
const requiredTag = "required"

/*FormatValidationErrors 处理参数校验错误
 * err 错误信息
 * descriptions 需要展示定制化错误信息的字段
 */
func FormatValidationErrors(
	err error, descriptions map[string]map[string]string,
) ([]FormatValidationError, error) {
	formatErrors := []FormatValidationError{}
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			tagName := fieldError.Tag()
			reason := fmt.Sprintf(defaultReason, fieldName)

			if tagName == requiredTag {
				reason = fmt.Sprintf(defaultRequiredReason, fieldName)
			} else {
				if descriptions != nil {
					if _, exist := descriptions[fieldName]; exist {
						if _, exist := descriptions[fieldName][tagName]; exist {
							reason = descriptions[fieldName][tagName]
						}
					}
				}
			}

			formatErrors = append(formatErrors, FormatValidationError{
				Field:  fieldName,
				Reason: reason,
			})
		}
	} else {
		return formatErrors, errors.New("not a valid validation error")
	}

	return formatErrors, nil
}
