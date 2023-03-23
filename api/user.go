package api

import (
	"BabyBus/config"
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
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

func GetUserInfo(ctx *gin.Context) {
	//封装JWT解析出来ID
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	if err := service.GetIdFromToken(user); err != nil {
		log.Printf("未从用户中成功获取用户id:%s\n", err)
		tool.Failure(400, "未成功从用户token中获取用户id", ctx)
		return
	}
	if err := service.GetUserInfo(user); err != nil {
		tool.Failure(400, "未成功保存user", ctx)
		log.Printf("未成功保存user:%s", err)
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}
