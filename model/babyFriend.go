package model

import "gorm.io/gorm"

type BabyFriend struct {
	gorm.Model
	UserId   string
	FriendId string
	Status   int //0:表示未响应 1:表示绑定成功 2:表示拒绝绑定
}
