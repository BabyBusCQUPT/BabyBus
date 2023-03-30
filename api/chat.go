package api

import (
	"BabyBus/config"
	"BabyBus/model"
	"BabyBus/service"
	"BabyBus/tool"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var rwLocker sync.RWMutex

func CreateConn(ctx *gin.Context) {
	//前端请求绑定好友
	//参数：申请人id 被请求人id
	user := &model.User{}
	user.Token = ctx.GetHeader("token")
	if err := service.GetIdFromToken(user); err != nil {
		log.Printf("从token中获取id失败:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("websocket upgrade error:%s\n", err)
		tool.Failure(500, "服务器错误", ctx)
		return
	}
	node := &model.Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 100),
	}
	if _, ok := config.ClientMap[user.OpenId]; ok {
		config.ClientMap[user.OpenId].Conn.Close()
		delete(config.ClientMap, user.OpenId)
	}

	rwLocker.Lock()
	config.ClientMap[user.OpenId] = node
	rwLocker.Unlock()

	//处理接收消息
	go func() {
		if err = service.RecProc(node, user.OpenId); err != nil {
			log.Printf("websocket receive error:%s\n", err)
		}
	}()
	//处理发送消息
	go func() {
		if err = service.SendProc(node); err != nil {
			log.Printf("websocket send error:%s\n", err)
		}
	}()
}
