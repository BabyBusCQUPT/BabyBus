package api

import (
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func GetBusScore(ctx *gin.Context) {
	busId := ctx.PostForm("busId")
	if err := tool.IsValid(busId); err != nil {
		tool.Failure(400, "巴士id为空", ctx)
		log.Printf("巴士id为空:%s", err)
	}
	sumScore, sumBaby, err := service.GetBusScore(busId)
	if err != nil {
		tool.Failure(500, "未从数据库中获取总评分，总人数", ctx)
		log.Printf("未从数据库中获取总评分，总人数:%s", err)
		return
	}
	ctx.JSON(200, gin.H{
		"sumScore": sumScore,
		"sumBaby":  sumBaby,
	})
}
