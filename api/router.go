package api

import (
	"BabyBus/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.Default()

	engine.Use(middleware.Cors)

	engine.POST("/user/register", Register)

	userGroup := engine.Group("/user")
	{
		userGroup.Use(middleware.TokenAuth)
		userGroup.DELETE("/delete", logOut)               //退出登录
		userGroup.POST("/update", Update)                 //更新用户信息
		userGroup.GET("/getPersonalInfo", GetUserInfo)    //获取用户基本信息
		userGroup.POST("/surroundings", UserSurroundings) //获取用户周边站点
	}

	babyBusGroup := engine.Group("/babyBus")
	{
		babyBusGroup.Use(middleware.TokenAuth)
		babyBusGroup.POST("/scoreBus", ScoreBus)           //用户打分
		babyBusGroup.GET("/singleScore", GetPersonalScore) //展示个人对巴士评分
	}

	friendsGroup := engine.Group("friends")
	{
		friendsGroup.Use(middleware.TokenAuth)
		wsGroup := friendsGroup.Group("/ws", CreateConn)
		{
			wsGroup.POST("/bindFriend", BindFriend) //绑定朋友 (站内消息绑定好友)
			wsGroup.POST("/addFriend", AddFriend)   //添加好友（通过ws发起请求）
			wsGroup.POST("/agree", Agree)           //ws同意绑定好友，好友添加成功
			wsGroup.POST("/reject", Reject)         //ws拒绝绑定好友
			wsGroup.POST("/share", Share)           //位置共享
		}
		friendsGroup.POST("/findFriend", DeriveFriend) //模糊搜索朋友
		friendsGroup.GET("/bindRecord", GetFriends)    //获取绑定好友及记录
	}

	busGroup := engine.Group("/bus")
	{
		busGroup.GET("/totalAverage", GetBusScore)       //展示巴士均分
		busGroup.POST("/fuzzyStation", FuzzyStation)     //模糊搜索站点
		busGroup.POST("/stationDetails", StationDetails) //获取站点详情信息
		busGroup.GET("/hotStations", HotStations)        //热搜站点
	}

	messageGroup := engine.Group("/message")
	{
		messageGroup.GET("/list", ListMsg)                 //罗列所有事件
		messageGroup.POST("/messageDetail", MessageDetail) //获取信息详情
	}

}
