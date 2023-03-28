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
	user := &model.User{}
	friendId := ctx.PostForm("friendId")
	err := tool.IsValid(friendId)
	if err != nil {
		tool.Failure(400, "绑定朋友失败：朋友id为空", ctx)
		return
	}
	friend := &model.User{OpenId: friendId}
	if err = service.GetUserInfo(friend); err != nil {
		log.Printf("未在表中查询到朋友，朋友未注册:%s\n", err)
		tool.Failure(400, "未在表中查询到朋友，朋友未注册", ctx)
		return
	}
	user.Token = ctx.GetHeader("token")
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("未从用户token中成功获取用户id:%s\n", err)
		tool.Failure(400, "未成功从用户token中获取用户id", ctx)
		return
	}
	if err = service.BindFriend(user.OpenId, friendId); err != nil {
		if err == config.AddHimself {
			tool.Failure(400, config.AddHimself.Error(), ctx)
			return
		} else if err == config.RepeatedAdd {
			tool.Failure(400, config.RepeatedAdd.Error(), ctx)
			return
		}
		log.Printf("绑定朋友id失败a:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"userId":   user.OpenId,
		"friendId": friendId,
		"context":  friend.Nickname + "请求与您绑定好友关系",
	})
}

func GetFriends(ctx *gin.Context) {
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	if err := service.GetIdFromToken(user); err != nil {
		log.Printf("未从用户中成功获取用户id:%s\n", err)
		tool.Failure(400, "未成功从用户token中获取用户id", ctx)
		return
	}
	openId := user.OpenId
	babyFriends := make([]*model.BabyFriend, 3, 6)
	if err := service.GetUserFriend(openId, babyFriends); err != nil {
		log.Printf("未从好友表中查询到好友:%s\n", err)
		tool.Failure(400, "未从好友表中查询到好友", ctx)
		return
	}
	type friendsInfo struct {
		Image string
		Date  time.Time
	}
	friendsInfos := make([]*friendsInfo, 3, 6)
	for i := range babyFriends {
		friendsInfos[i].Date = babyFriends[i].CreatedAt
		user := &model.User{}
		user.OpenId = babyFriends[i].FriendId
		if err := service.GetUserInfo(user); err != nil {
			log.Printf("未从用户表中查询到该用户的好友:%s\n", err)
			tool.Failure(400, "未从用户表中查询到该用户的好友", ctx)
			return
		}
		friendsInfos[i].Image = user.Image
	}
	ctx.JSON(http.StatusOK, gin.H{
		"friendsInfos": friendsInfos,
	})
}
