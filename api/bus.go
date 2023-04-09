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

func StationDetails(ctx *gin.Context) {
	var err error
	stationName := ctx.PostForm("stationName")
	if err = tool.IsValid(stationName); err != nil {
		tool.Failure(400, "缺失站点名称：关键字为空", ctx)
		return
	}
	station, err := service.GetStationDetails(stationName)
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("查询站点详情信息失败:%s\n", err)
		return
	}
	if err = service.StationsScoreIncr(1, stationName); err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("增加用户查询次数失败:%s\n", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"stationDInfo": station,
	})
}

func HotStations(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"hot":  service.GetHot(),
	})
}
