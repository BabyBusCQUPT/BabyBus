package api

import (
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MessageDetail(ctx *gin.Context) {
	user := &model.User{}
	message := &model.Message{}
	message.PostId = ctx.PostForm("postId")
	user.Token = ctx.PostForm("token")
	if err := tool.IsValid(message.PostId); err != nil {
		tool.Failure(400, "缺失必要参数:缺失发信人id", ctx)
		return
	}
	if err := service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	message.ReceiveId = user.OpenId
	if err := service.SelectMsgDetail(message); err != nil {
		log.Printf("获取事件详情失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"detail": message,
	})
}
