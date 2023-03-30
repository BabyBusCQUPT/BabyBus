package model

import (
	"github.com/gorilla/websocket"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
}
