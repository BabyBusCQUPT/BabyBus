package config

import "BabyBus/model"

var (
	ClientMap = make(map[string]*model.Node, 0)
	Yes       = "1" //同意好友申请
	No        = "0" //拒绝好友申请
	Rejected  = "已拒绝你的好友绑定申请"
	Accepted  = "已通过你的好友绑定申请"
)
