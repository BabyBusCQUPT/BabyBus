package api

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.Use()
		userGroup.POST("/register", Register)          //微信注册
		userGroup.DELETE("/delete", logOut)            //退出登录
		userGroup.POST("/update", Update)              //更新用户信息
		userGroup.GET("/getPersonalInfo", GetUserInfo) //获取用户基本信息
	}

	babyBusGroup := engine.Group("/babyBus")
	{

		babyBusGroup.POST("/scoreBus", ScoreBus)           //用户打分
		babyBusGroup.GET("/singleScore", GetPersonalScore) //展示个人对巴士评分
	}

	friendsGroup := engine.Group("friends")
	{
		friendsGroup.POST("/findFriend", DeriveFriend) //模糊搜索朋友
		friendsGroup.POST("/bingFriend", BindFriend)   //绑定朋友
		friendsGroup.POST("/addFriend", AddFriend)     //添加好友（通过ws发起请求）
		friendsGroup.GET("/bindRecord", GetFriends)    //获取绑定好友及记录
	}

	busGroup := engine.Group("/bus")
	{
		busGroup.GET("/totalAverage", GetBusScore) //展示巴士均分
	}

	messageGroup := engine.Group("/message")
	{
		messageGroup.GET("/list", ListMsg)                 //罗列所有事件
		messageGroup.POST("/messageDetail", MessageDetail) //获取信息详情
	}

}
