package dao

import "BabyBus/model"

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

/*
// AcceptFriend 绑定朋友成功
func AcceptFriend(userId string, friendId string) error {
	//status 表示连接成功
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ? and friend_id = ?", userId, friendId).Update("status", 1).Error; err != nil {
		return err
	}
	return nil
}

// RejectFriend 绑定朋友被拒绝
func RejectFriend(userId string, friendId string) error {
	if err := DB.Model(&model.BabyFriend{}).Where("user_id = ? and friend_id = ?", userId, friendId).Update("status", 2).Error; err != nil {
		return err
	}
	return nil
}
*/
