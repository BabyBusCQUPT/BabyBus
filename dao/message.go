package dao

import "BabyBus/model"

func SelectMessageDetail(message *model.Message) error {
	return DB.Model(&model.Message{}).Where("postId = ? AND receiveId = ?", message.PostId, message.ReceiveId).Find(&message).Error
}
