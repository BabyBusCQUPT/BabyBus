package middleware

import (
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(c *gin.Context) {
	user := &model.User{}
	WeChatConnection := &model.WeChatConnection{}
	user.Token = c.GetHeader("token")
	if user.Token == "" {
		tool.Failure(400, "未获得合理的token", c)
		return
	}
	tokenClaims, err := service.ParseToken(user.Token)
	if err != nil {
		tool.Failure(400, "未成功解析token", c)
		//没有判断token是否过期
		return
	}

	identify := tokenClaims.Identify
	err = json.Unmarshal([]byte(identify), &WeChatConnection)
	if err != nil {
		tool.Failure(400, "未解析到相应字段", c)
		return
	}
	openId := WeChatConnection.OpenId
	sessionKey := WeChatConnection.SessionKey

	user.OpenId = openId
	user.SessionKey = sessionKey

}
