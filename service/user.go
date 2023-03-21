package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func DeleteToken(user model.User) error {
	return dao.Delete(user)
}

// CountAllId 获取用户id（数据库主键）
func CountAllId() (id int64, err error) {
	return dao.CountAllId()
}

// SaveUser 注册时存储用户信息
func SaveUser(user *model.User) error {
	return dao.Save(user)
}

// UpdateUser 更新用户信息
func UpdateUser(user model.User) error {
	return dao.Update(user)
}

// SearchByKeyWords 模糊搜索昵称
func SearchByKeyWords(words string) ([]model.User, error) {
	return dao.Search(words)
}

func BindFriend(user *model.User) error {
	return dao.Save(user)
}
