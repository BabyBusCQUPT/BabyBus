package dao

import (
	"BabyBus/model"
)

func Delete(user model.User) error {
	if err := DB.Model(&user).Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func CountAllId() (id int64, err error) {
	if err := DB.Model(&model.User{}).Count(&id).Error; err != nil {
		return 0, err
	}
	return id + 1, nil
}
