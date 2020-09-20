package user

import (
	"domisep/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*CreateUserController 创建用户函数
 * params: c *gin.Context: 上下文
 */
func CreateUserController(c *gin.Context) {
	var userInfo UserObject
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid user information!",
		})
		return
	}

	db.DB.Save(&userInfo)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "User created successfully!",
		"resourceId": userInfo.ID,
	})
}
