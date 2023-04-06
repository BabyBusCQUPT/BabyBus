package dao

import (
	"BabyBus/config"
	"BabyBus/model"
)

// GetUserFriends 获取绑定的所有朋友
func GetUserFriends(openId string) (babyFriend []model.BabyFriend, err error) {
	//var babyFriend []model.BabyFriend
	if err = DB.Model(&model.BabyFriend{}).Where("user_id = ? AND status = ?", openId, config.Accept).Find(&babyFriend).Error; err != nil {
		return nil, err
	}
	return babyFriend, nil
}

// BindFriend 请求绑定朋友，状态显示为在绑定中
func BindFriend(userId string, friendId string) error {
	if userId == friendId {
		return config.AddHimself
	}
	babyFriend := &model.BabyFriend{
		UserId:   userId,
		FriendId: friendId,
	}
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ? and friend_id = ?", userId, friendId).Find(&babyFriend).Error; err != nil {
		return err
	}
	if babyFriend.ID != 0 {
		return config.RepeatedAdd
	}
	if err := DB.Model(&model.BabyFriend{}).Save(babyFriend).Error; err != nil {
		return err
	}
	babyFriend.UserId = friendId
	babyFriend.FriendId = userId
	if err := DB.Model(&model.BabyFriend{}).Save(babyFriend).Error; err != nil {
		return err
	}
	return nil
}

func CountFriend(openId string) (int64, error) {
	var count int64
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ?", openId).Count(&count).Error; err != nil {
		return config.InvalidParameter, err
	}
	return count, nil
}

// AcceptFriend 绑定朋友成功
func AcceptFriend(userId string, friendId string) error {
	//status 表示连接成功
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ? and friend_id = ?", userId, friendId).Update("status", 1).Error; err != nil {
		return err
	}
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ? and friend_id = ?", friendId, userId).Update("status", 1).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFriend 绑定朋友被拒绝
func DeleteFriend(userId string, friendId string) error {
	if err := DB.Model(&model.BabyFriend{}).Delete(model.BabyFriend{UserId: userId, FriendId: friendId}).Error; err != nil {
		return err
	}
	if err := DB.Model(&model.BabyFriend{}).Delete(model.BabyFriend{UserId: friendId, FriendId: userId}).Error; err != nil {
		return err
	}
	return nil
}
