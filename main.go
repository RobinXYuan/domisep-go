package main

import (
	"domisep/apps/auth"
	"domisep/apps/user"
	"domisep/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfigs()
}

func registerRouters(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		auth.RegisterRouters(v1)
		user.RegisterRouters(v1)
	}
}

func main() {
	engine := gin.Default()

	registerRouters(engine)

	engine.Run(":8080")
}
