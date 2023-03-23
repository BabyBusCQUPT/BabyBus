package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func GetUserFriend(openid string, babyFriend []*model.BabyFriend) error {
	return dao.GetUserFriend(openid, babyFriend)
}
