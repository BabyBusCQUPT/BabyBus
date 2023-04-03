package api

import (
	"BabyBus/config"
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// DeriveFriend 模糊搜索朋友
func DeriveFriend(ctx *gin.Context) {
	friendName := ctx.PostForm("friendName")
	err := tool.IsValid(friendName)
	if err != nil {
		tool.Failure(400, "查找朋友失败：昵称字段为空", ctx)
		return
	}
	friends, err := service.SearchByKeyWords(friendName)
	if err != nil {
		fmt.Printf("模糊搜索nickname失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"friends": friends,
	})
}

func BindFriend(ctx *gin.Context) {
	var err error
	friend := &model.User{}
	user := &model.User{}
	friend.OpenId = ctx.PostForm("friendId") //之前申请绑定好友的用户的id
	choice := ctx.PostForm("choice")
	user.Token = ctx.GetHeader("token") //目前处理站内消息的用户的id

	if err = tool.IsValid(friend.OpenId); err != nil {
		tool.Failure(400, "绑定朋友失败：朋友id为空", ctx)
		return
	}
	if err = tool.IsValid(choice); err != nil {
		tool.Failure(400, "绑定朋友失败：选择为空", ctx)
		return
	}

	if choice == config.No {
		//拒绝好友绑定
		//删除原来绑定记录
		if err = service.RejectFriend(friend.OpenId, user.OpenId); err != nil {
			tool.Failure(500, "服务器错误", ctx)
			log.Printf("站内消息绑定好友失败：删除绑定记录失败")
			return
		}
		service.SendMsg(friend.OpenId, config.Rejected)
		tool.Success("successfully rejected", ctx)
		return
	}

	if err = service.GetUserInfo(friend); err != nil {
		log.Printf("未在表中查询到朋友，朋友未注册:%s\n", err)
		tool.Failure(400, "未在表中查询到朋友，朋友未注册", ctx)
		return
	}
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("未从用户token中成功获取用户id:%s\n", err)
		tool.Failure(400, "未成功从用户token中获取用户id", ctx)
		return
	}
	num, err := service.CountFriend(user.OpenId)
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("统计当前好友数量失败:%s\n", err)
		return
	}
	if num > config.MaxFriendNum {
		tool.Failure(400, "绑定好友数已达上限", ctx)
		return
	}
	service.SendMsg(friend.OpenId, config.Accepted)
	ctx.JSON(200, gin.H{
		"code": http.StatusOK,
		"info": "绑定成功",
	})
}

// GetFriends 获取已经绑定的所有朋友
func GetFriends(ctx *gin.Context) {
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	if err := service.GetIdFromToken(user); err != nil {
		log.Printf("未从用户中成功获取用户id:%s\n", err)
		tool.Failure(400, "未成功从用户token中获取用户id", ctx)
		return
	}
	babyFriends, err := service.GetUserFriends(user.OpenId)
	if err != nil {
		log.Printf("未从好友表中查询到好友:%s\n", err)
		tool.Failure(400, "未从好友表中查询到好友", ctx)
		return
	}
	type friendsInfo struct {
		Image    string
		NikeName string
		Date     time.Time
	}
	friendsInfos := make([]*friendsInfo, 3, 6)
	for i := range babyFriends {
		friendsInfos[i].Date = babyFriends[i].CreatedAt
		friend := &model.User{}
		friend.OpenId = babyFriends[i].FriendId
		if err = service.GetUserInfo(friend); err != nil {
			log.Printf("未从用户表中查询到该用户的好友:%s\n", err)
			tool.Failure(400, "未从用户表中查询到该用户的好友", ctx)
			return
		}
		friendsInfos[i].Image = friend.Image
		friendsInfos[i].NikeName = friend.Nickname
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"friendsInfos": friendsInfos,
	})
}

func AddFriend(ctx *gin.Context) {
	var err error
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	babyFriend := &model.BabyFriend{
		UserId: user.OpenId,
	}
	babyFriend.FriendId = ctx.PostForm("friendId")
	if err = tool.IsValid(babyFriend.FriendId); err != nil {
		tool.Failure(400, "缺失必要参数：缺失friendId", ctx)
		return
	}
	if err = service.BindFriend(babyFriend.UserId, babyFriend.FriendId); err != nil {
		log.Printf("绑定好友失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	service.SendMsg(user.OpenId, user.Nickname+"邀请您绑定好友关系")
}
