package model

import (
	"time"
)

type Message struct {
	PostId    int
	ReceiveId int
	PostTime  time.Time
	Detail    string
}
