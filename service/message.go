package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func SelectMsgDetail(message *model.Message) error {
	return dao.SelectMessageDetail(message)
}

func ListMsg(message model.List) ([]model.List, error) {
	return dao.ListMsg(message)
}
