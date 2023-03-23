package dao

import "BabyBus/model"

func GetUserFriend(openId string, babyfriend []*model.BabyFriend) error {
	if err := DB.Model(&model.BabyFriend{}).Where("openid = ?", openId).Find(&babyfriend).Error; err != nil {
		return err
	}
	return nil
}
