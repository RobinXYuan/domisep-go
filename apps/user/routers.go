package user

import "github.com/gin-gonic/gin"

//RegisterRouters 注册路由
func RegisterRouters(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.POST("/", CreateUserController)
	}
}
