package api

import (
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
	identify := "{\"OpenId\":\"" + weChatConnection.OpenId + "\"," +
		"\"SessionKey\":\"" + weChatConnection.SessionKey + "\"}"
	token, err := service.CreateToken(identify)
	if err != nil {
		log.Printf("生成token失败")
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	tool.Success(token, ctx)
}
