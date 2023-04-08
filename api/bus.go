package api

import (
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func FuzzyStation(ctx *gin.Context) {
	var err error
	words := ctx.PostForm("keyWords")
	if err = tool.IsValid(words); err != nil {
		tool.Failure(400, "缺失必要参数：关键字为空", ctx)
		return
	}
	stations, err := service.SelectStations(words)
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("模糊搜索站点失败:%s\n", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"stations": stations,
	})
}
