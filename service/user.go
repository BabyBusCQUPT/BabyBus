package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func DeleteToken(user model.User) error {
	return dao.Delete(user)
}

func CountAllId() (id int64, err error) {
	return dao.CountAllId()
}
