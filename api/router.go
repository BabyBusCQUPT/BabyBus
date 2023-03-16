package api

import "github.com/gin-gonic/gin"

func Init() {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.Use()
		userGroup.POST("/register", Register) //微信注册
	}

}
