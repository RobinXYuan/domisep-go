package auth

import "github.com/gin-gonic/gin"

//RegisterRouters 注册路由
func RegisterRouters(group *gin.RouterGroup) {
	authRouter := group.Group("/auth")
	{
		authRouter.POST("/login", Login)
		authRouter.POST("/register", Register)
	}
}
