package service

import "BabyBus/model"

func SelectMsgDetail(message *model.Message) error {
	return dao.SelectMessageDetail(message)
}
