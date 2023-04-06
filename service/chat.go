package service

import (
	"BabyBus/config"
	"BabyBus/model"
	"errors"
	"github.com/gorilla/websocket"
)

// SendProc 发送消息，即将消息写入管道，供前端取出来
func SendProc(nodeUser *model.Node, nodeFriend *model.Node) error {
	for {
		select {
		case data := <-nodeUser.DataQueue:
			err := nodeUser.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return errors.New("发送信息错误，将信息写入管道错误" + err.Error())
			}
		case data := <-nodeFriend.DataQueue:
			err := nodeFriend.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return errors.New("发送信息错误，将信息写入管道错误" + err.Error())
			}
		}
	}
}

// RecProc 接收消息，即将消息从管道里面读出来，转交给另一个人
func RecProc(nodeUser *model.Node, userId string, nodeFriend *model.Node, friendId string) error {
	var d []byte
	for {
		select {
		case nodeUser.DataQueue <- d:
			_, data, err := nodeUser.Conn.ReadMessage()
			if err != nil {
				if err.Error() == "websocket: close 1001 (going away)" {
					return errors.New("websocket连接断开")
				}
				return errors.New("in recProc read message error: " + err.Error())
			}
			SendMsg(userId, string(data))
		case nodeFriend.DataQueue <- d:
			_, data, err := nodeFriend.Conn.ReadMessage()
			if err != nil {
				if err.Error() == "websocket: close 1001 (going away)" {
					return errors.New("websocket连接断开")
				}
				return errors.New("in recProc read message error: " + err.Error())
			}
			SendMsg(friendId, string(data))
		}

	}
}

func Close(id string) error {
	return config.ClientMap[id].Conn.Close()
}
