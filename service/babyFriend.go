package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func GetUserFriend(openid string, babyFriend []*model.BabyFriend) error {
	return dao.GetUserFriend(openid, babyFriend)
}

/*
// AcceptFriend 绑定朋友被同意
func AcceptFriend(userId string, friendId string) error {
	return dao.AcceptFriend(userId, friendId)
}

// RejectFriend 绑定朋友被拒绝
func RejectFriend(userId string, friendId string) error {
	return dao.RejectFriend(userId, friendId)
}
*/
