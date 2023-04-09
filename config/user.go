package config

import "time"

const (
	InvalidParameter = -1 //参数无效返回值
	MaxFriendNum     = 6  //允许最大好友数

	Pending = 0 //等待回应
	Accept  = 1 //同意绑定
	Reject  = 2 //拒绝绑定

	LimitedRequest = 5               //十分钟之类可以请求的次数：防止恶意请求
	Incr           = 1               //请求次数偏移量
	Expire         = 3 * time.Minute //请求过期时间
)
