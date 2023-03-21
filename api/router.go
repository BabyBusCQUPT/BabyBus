package api

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.Use()
		userGroup.POST("/register", Register)       //微信注册
		userGroup.DELETE("/delete", logOut)         //退出登录
		userGroup.POST("/update", Update)           //更新用户信息
		userGroup.POST("/scoreBus", ScoreBus)       //用户打分
		userGroup.POST("/findFriend", DeriveFriend) //模糊搜索朋友
		userGroup.POST("/bingFriend", BindFriend)   //绑定朋友
	}

}
