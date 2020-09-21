package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const domisepSchema = "domisep"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(domisepSchema):]
		token, err := utils.ValidateToken(tokenString)
		if token.Valid {

		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
