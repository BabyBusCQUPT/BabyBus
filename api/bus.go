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

	//防止恶意刷新导致热榜失真的问题
	//获取ip地址
	ip := ctx.ClientIP()
	//查看当前是否过量访问
	if err = service.CheckLimit(ip); err != nil {
		if err.Error() == config.TooManyRequests.Error() {
			tool.Failure(http.StatusTooManyRequests, config.TooManyRequests.Error(), ctx)
			return
		}
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	//未过量访问则更新访问状态
	if err = service.IpRefresh(ip); err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("更新ip访问错误:%s\n", err)
		return
	}

	station, err := service.GetStationDetails(stationName)
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("查询站点详情信息失败:%s\n", err)
		return
	}
	if err = service.StationsScoreIncr(config.Incr, stationName); err != nil {
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
