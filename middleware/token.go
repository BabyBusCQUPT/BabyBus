package middleware

import (
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func TokenAuth(ctx *gin.Context) {
	user := model.User{}
	user.Token = ctx.GetHeader("token")
	if err := tool.IsValid(user.Token); err != nil {
		tool.Failure(400, "token为空", ctx)
		return
	}
	tokenClaims, err := service.ParseToken(user.Token)
	if err != nil {
		if tokenClaims.ExpireTime.Before(time.Now()) {
			tool.Failure(400, "token已经过期", ctx)
			return
		}
		tool.Failure(500, "解析token失败", ctx)
		log.Println("解析token失败")
		return
	}
	if err = service.ParseTokenIdentify(&user, tokenClaims); err != nil {
		tool.Failure(500, "解析openId和SessionKey出错", ctx)
		log.Printf("解析openId和SessionKey出错%s\n", err)
		return
	}

	token, err := service.FindOneWithOpenIdAndSessionKey(user)
	if err != nil {
		tool.Failure(400, "用户不存在该token，查询失败", ctx)
		log.Printf("用户不存在该token，查询失败:%s\n", err)
		return
	}
	if token != user.Token {
		tool.Failure(400, "传入token与用户不相对应", ctx)
		log.Printf("用户token与传入token不匹配")
		return
	}
}
