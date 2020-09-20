package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//AuthSessionMiddleware session 认证中间件
func AuthSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized user!",
			})
			c.Abort()
			return
		}

		c.Set("userId", sessionValue.(uint))

		c.Next()
		return
	}
}
