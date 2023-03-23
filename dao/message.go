package dao

import "BabyBus/model"

func SelectMessageDetail(message *model.Message) error {
	return DB.Model(&model.Message{}).Where("postId = ? AND receiveId = ? AND id = ?", message.PostId, message.ReceiveId, message.ID).First(&message).Error
}
