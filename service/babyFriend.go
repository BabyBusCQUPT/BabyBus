package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

// GetUserFriends 获取用户所有的朋友来进行展示
func GetUserFriends(openid string) (babyFriend []model.BabyFriend, err error) {
	return dao.GetUserFriends(openid)
}

// BindFriend 请求绑定朋友
func BindFriend(userId string, friendId string) error {
	return dao.BindFriend(userId, friendId)
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
