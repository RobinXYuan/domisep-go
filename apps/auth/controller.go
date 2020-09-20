package auth

import (
	"domisep/constants"
	"domisep/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var registerValidationDescription = map[string]map[string]string{
	"name": {
		"min": "The length of user name is at least 4 characters!",
		"max": "The length of user name is at most 80 characters!",
	},
	"password": {
		"password": `The password must contain at least 8 characters and at least a capital character, a number and a special character in '@!_-=:*'`,
	},
	"confirmPassword": {
		"eqfield": "The value of confirmPassword must be equal to password field!",
	},
}

//Login 登录函数
func Login(c *gin.Context) {
	loginParams := &LoginParams{}
	pass, reasons := utils.ValidateParameters(c, loginParams, constants.JSONTypeParam, nil)
	if !pass {
		c.AbortWithStatusJSON(http.StatusBadRequest, &constants.ErrorReturn{
			Status:  http.StatusBadRequest,
			Message: reasons["message"].(string),
			Errors:  reasons["errors"],
		})
	}

	// if hasSession := utils.HasSession(c); hasSession == true {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  http.StatusOK,
	// 		"message": fmt.Sprintf("User %s has already been logged in!", loginUserObj.Name),
	// 	})
	// 	return
	// }

	// // userObj := user.GetUserDetailByName(loginUserObj.Name)

	// if err := utils.ComparePassword(userObj.Password, loginUserObj.Password); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status":  http.StatusUnauthorized,
	// 		"message": "Incorrect password or username!",
	// 	})
	// 	return
	// }

	// utils.SaveAuthSession(c, userObj.ID)
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  http.StatusOK,
	// 	"message": "Login successfully!",
	// })
	return
}

//Register 注册函数
func Register(c *gin.Context) {
	registerParams := &RegisterParams{}
	pass, reasons := utils.ValidateParameters(
		c, registerParams,
		constants.JSONTypeParam, registerValidationDescription,
	)
	if !pass {
		c.AbortWithStatusJSON(http.StatusBadRequest, &constants.ErrorReturn{
			Status:  http.StatusBadRequest,
			Message: reasons["message"].(string),
			Errors:  reasons["errors"],
		})
	}
}
