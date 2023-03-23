package dao

import "BabyBus/model"

func SelectMessageDetail(message *model.Message) error {
	return DB.Model(&model.Message{}).Where("postId = ? AND receiveId = ? AND id = ?", message.PostId, message.ReceiveId, message.ID).First(&message).Error
}

func ListMsg(message model.List) (list []model.List, err error) {
	result := DB.Model(&model.Message{}).Where("postId = ? AND receiveId = ?", message.PostId, message.ReceiveId).Find(&list)
	return list, result.Error
}
