package model

import "gorm.io/gorm"

type BabyFriend struct {
	gorm.Model
	UserId   string
	FriendId string
}
