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
)

func Register(ctx *gin.Context) {
	applet, err := service.ParseAppletConfig()
	if err != nil {
		log.Printf("解析applet配置文件失败：%s", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}

	code := ctx.PostForm("code")
	weChatConnection, err := service.ConnectWeChatApi(applet, code)
	if err != nil {
		log.Printf("连接微信服务失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}

	all, err := service.CountAllId()
	if err != nil {
		log.Printf("获取用户id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	//创建token的identify字段
	identify := "{\"OpenId\":\"" + weChatConnection.OpenId + "\"," +
		"\"SessionKey\":\"" + weChatConnection.SessionKey + "\"," +
		"{\"Id\":\"" + string(all) + "\"}"
	token, err := service.CreateToken(identify)
	if err != nil {
		log.Printf("生成token失败")
		tool.Failure(500, "服务器错误", ctx)
		return
	}

	user := &model.User{}
	user.ID = uint(all)
	user.OpenId = weChatConnection.OpenId
	user.SessionKey = weChatConnection.SessionKey
	user.Token = token
	if err = service.SaveUser(user); err != nil {
		log.Printf("未成功存储用户信息:%s\n", err)
		tool.Failure(400, "未成功存储用户信息", ctx)
		return
	}

	tool.Success(token, ctx)
}

func logOut(ctx *gin.Context) {
	user := model.User{}
	user.Token = ctx.GetHeader("token")
	err := service.DeleteToken(user)
	if err != nil {
		log.Printf("软删除消息失败：%s", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	tool.Success("成功退出登录", ctx)
}

func Update(ctx *gin.Context) {
	user := model.User{}
	age := ctx.PostForm("age")
	gender := ctx.PostForm("gender")
	image := ctx.PostForm("image")
	nickname := ctx.PostForm("nickname")
	//判断是否为空
	//转化类型后保存
	var err error
	if user.Age, err = tool.IsValidAndTrans(age); err != nil {
		if err == config.InvalidParameterErr {
			tool.Failure(400, "用户修改信息错误：字符串为空", ctx)
			return
		}
		log.Printf("转化int错误：%s\n", err)
		tool.Failure(400, "用户修改信息错误：转int错误", ctx)
		return
	}

	g, err := tool.IsValidAndTrans(gender)
	if err != nil {
		if err == config.InvalidParameterErr {
			tool.Failure(400, "用户修改信息错误：字符串为空", ctx)
			return
		}
		log.Printf("转化int错误：%s\n", err)
		tool.Failure(400, "用户修改信息错误：转int错误", ctx)
		return
	}
	user.Gender = byte(g)

	if err = tool.IsValid(image); err != nil {
		tool.Failure(400, "用户修改信息错误：传入头像为空", ctx)
		return
	}
	user.Image = image

	if err = tool.IsValid(nickname); err != nil {
		tool.Failure(400, "用户修改信息错误：传入昵称为空", ctx)
		return
	}
	user.Nickname = nickname

	if err = service.UpdateUser(user); err != nil {
		tool.Failure(400, "用户修改信息错误", ctx)
		return
	}
}

func ScoreBus(ctx *gin.Context) {
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	tokenClaims, err := service.ParseToken(user.Token)
	if err != nil {
		log.Printf("解析token失败：%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	err = service.ParseTokenIdentify(user, tokenClaims)
	if err != nil {
		log.Printf("获取token内用户信息失败：%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	if err = service.SaveUser(user); err != nil {
		log.Printf("存储用户打分失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	tool.Success("打分成功", ctx)
}

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

// BindFriend 绑定朋友
func BindFriend(ctx *gin.Context) {
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	id := ctx.PostForm("friend")
	i, err := tool.IsValidAndTrans(id)
	if err != nil {
		if err == config.InvalidParameterErr {
			tool.Failure(400, "绑定朋友失败：朋友id为空", ctx)
			return
		}
		log.Printf("string转int失败：%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	user.Friend = uint(i)
	if err = service.BindFriend(user); err != nil {
		log.Printf("绑定朋友id失败a:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	tool.Success("成功绑定朋友", ctx)
	return
}
