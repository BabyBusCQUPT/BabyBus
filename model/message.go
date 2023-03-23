package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	PostId    string
	ReceiveId string
	PostTime  time.Time
	Detail    string
}

type List struct {
	gorm.Model
	PostId    string
	ReceiveId string
	PostTime  time.Time
}
