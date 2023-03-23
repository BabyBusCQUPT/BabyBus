package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	PostId    int
	ReceiveId int
	PostTime  time.Time
	Detail    string
}
