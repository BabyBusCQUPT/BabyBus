package config

import "errors"

var (
	InvalidParameterErr = errors.New("invalid parameters")
	RepeatedAdd         = errors.New("重复添加相同好友")
	AddHimself          = errors.New("添加自己为好友")
)
