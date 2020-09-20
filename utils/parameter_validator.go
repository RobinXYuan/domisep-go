package utils

import (
	"domisep/constants"

	"github.com/gin-gonic/gin"
)

//ValidateParameters 参数校验函数
func ValidateParameters(
	context *gin.Context, params interface{}, paramsType constants.ParameterType,
	descriptions map[string]map[string]string,
) (bool, map[string]interface{}) {
	switch paramsType {
	case constants.JSONTypeParam:
		return validateJSONParameters(context, params, descriptions)
	case constants.FormTypeParam:
		return validateFormParameters(context, params, descriptions)
	default:
		return false, nil
	}
}

func validateJSONParameters(
	c *gin.Context, params interface{},
	descriptions map[string]map[string]string,
) (bool, map[string]interface{}) {
	var validateResult = make(map[string]interface{})
	if err := c.ShouldBindJSON(params); err != nil {
		formatError, err := FormatValidationErrors(err, descriptions)
		if err != nil {
			validateResult = map[string]interface{}{
				"message": "Invalid input!",
				"errors":  err,
			}
			return false, validateResult
		}

		validateResult = map[string]interface{}{
			"message": "Invalid input!",
			"errors":  formatError,
		}
		return false, validateResult
	}

	return true, validateResult
}

func validateFormParameters(
	c *gin.Context, params interface{},
	descriptions map[string]map[string]string,
) (bool, map[string]interface{}) {
	return true, make(map[string]interface{})
}
