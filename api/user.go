package api

import (
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
		log.Printf("连接微信服务失败")
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	all, err := service.CountAllId()
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	identify := "{\"OpenId\":\"" + weChatConnection.OpenId + "\"," +
		"\"SessionKey\":\"" + weChatConnection.SessionKey + "\"," +
		"{\"Id\":\"" + string(all) + "\"}"
	token, err := service.CreateToken(identify)
	if err != nil {
		log.Printf("生成token失败")
		tool.Failure(500, "服务器错误", ctx)
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
