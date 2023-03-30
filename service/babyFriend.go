package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

// GetUserFriends 获取用户所有的朋友来进行展示
func GetUserFriends(openid string) (babyFriend []model.BabyFriend, err error) {
	return dao.GetUserFriends(openid)
}

// BindFriend 请求绑定朋友
func BindFriend(userId string, friendId string) error {
	friend := model.User{
		OpenId: friendId,
	}
	if err := dao.GetUserInfo(&friend); err != nil {
		return err
	}
	if friend.ID == 0 {
		return errors.New("用户邀请绑定的好友不存在")
	}
	user := model.User{
		OpenId: userId,
	}
	if err := dao.GetUserInfo(&user); err != nil {
		return err
	}
	if err := dao.BindFriend(userId, friendId); err != nil {
		return err
	}
	SendMsg(userId, user.Nickname+"邀请您绑定好友关系")
	return nil
}

func CountFriend(openId string) (int64, error) {
	return dao.CountFriend(openId)
}

// AcceptFriend 绑定朋友被同意
func AcceptFriend(userId string, friendId string) error {
	if err := dao.AcceptFriend(userId, friendId); err != nil {
		return err
	}
	SendMsg(userId, "绑定成功")
	return nil
}

// RejectFriend 绑定朋友被拒绝
func RejectFriend(userId string, friendId string) error {
	if err := dao.RejectFriend(userId, friendId); err != nil {
		return err
	}
	SendMsg(userId, "绑定好友失败")
	return nil
}

func SendMsg(userId string, msg string) {
	rwLocker.RLock()
	defer rwLocker.RUnlock()
	node := config.ClientMap[userId]
	node.DataQueue <- []byte(msg)
}
