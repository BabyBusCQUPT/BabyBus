package api

import (
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func MessageDetail(ctx *gin.Context) {
	var err error
	var id int
	user := &model.User{}
	message := &model.Message{}
	message.PostId = ctx.PostForm("postId")
	messageId := ctx.PostForm("msgId")
	user.Token = ctx.PostForm("token")
	if err := tool.IsValid(message.PostId); err != nil {
		tool.Failure(400, "缺失必要参数:缺失发信人id", ctx)
		return
	}
	if id, err = tool.IsValidAndTrans(messageId); err != nil {
		tool.Failure(400, "缺失必要参数:信息id", ctx)
		return
	}
	message.ID = uint(id)
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	message.ReceiveId = user.OpenId
	if err = service.SelectMsgDetail(message); err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("未查询到相关信息:%s\n", err)
			tool.Failure(400, "事件id错误", ctx)
			return
		}
		log.Printf("获取事件详情失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"detail": message,
	})
}

func ListMsg(ctx *gin.Context) {
	user := &model.User{}
	message := model.List{}
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
	list, err := service.ListMsg(message)
	if err != nil {
		log.Printf("查询所有信息失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"list": list,
	})
}
