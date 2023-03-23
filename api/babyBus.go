package api

import (
	"BabyBus/config"
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ScoreBus(ctx *gin.Context) {
	var err error
	user := &model.User{}
	scoreBus := &model.BabyBus{
		BabyId: user.OpenId,
	}
	user.Token = ctx.GetHeader("token")
	busId := ctx.PostForm("busId")
	score := ctx.PostForm("score")
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	if err = tool.IsValid(score); err != nil {
		tool.Failure(400, "缺失必要参数:缺失分数字段", ctx)
		return
	}
	if scoreBus.Score, err = tool.StringToFloat(score); err != nil {
		log.Printf("string转化为float失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	if scoreBus.BusId, err = tool.IsValidAndTrans(busId); err != nil {
		if err == config.InvalidParameterErr {
			tool.Failure(400, "缺失必要参数：busId为空", ctx)
			return
		}
		log.Printf("string to int fail:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}

	//保存个人评分记录
	if err = service.SaveScore(scoreBus); err != nil {
		log.Printf("保存打分失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}

	tool.Success("打分成功", ctx)
}

func GetPersonalScore(ctx *gin.Context) {
	var err error
	user := &model.User{}
	personalScore := &model.BabyBus{}
	user.Token = ctx.GetHeader("token")
	busId := ctx.PostForm("busId")
	if err = service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取openId失败：%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	if personalScore.BusId, err = tool.IsValidAndTrans(busId); err != nil {
		if err == config.InvalidParameterErr {
			tool.Failure(400, "缺失必要参数：busId为空", ctx)
			return
		}
		log.Printf("string to int fail:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	personalScore.BabyId = user.OpenId
	scores, err := service.GetPersonalScore(personalScore)
	if err != nil {
		tool.Failure(500, "服务器错误", ctx)
		log.Printf("查询用户对巴士评分失败:%s\n", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"scores": scores,
	})
}

func GetBusScore(ctx *gin.Context) {

}
