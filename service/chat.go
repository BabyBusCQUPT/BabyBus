package service

import (
	"BabyBus/model"
	"errors"
	"github.com/gorilla/websocket"
)

// SendProc 发送消息，即将消息写入管道，供前端取出来
func SendProc(node *model.Node) error {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return errors.New("发送信息错误，将信息写入管道错误" + err.Error())
			}
		}
	}
}

// RecProc 接收消息，即将消息从管道里面读出来，转交给另一个人
func RecProc(node *model.Node, userId string) error {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			if err.Error() == "websocket: close 1001 (going away)" {
				return errors.New("websocket连接断开")
			}
			return errors.New("in recvProc read message error: " + err.Error())
		}
		SendMsg(userId, string(data))
	}
}

func Close(id string) error {
	return config.ClientMap[id].Conn.Close()
}
