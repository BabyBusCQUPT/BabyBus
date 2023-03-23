package tool

import (
	"BabyBus/config"
	"strconv"
)

// IsValidAndTrans 判断字符串是否为空并且转化为int类型
func IsValidAndTrans(s string) (int, error) {
	if s == "" {
		return config.InvalidParameter, config.InvalidParameterErr
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return config.InvalidParameter, err
	}
	return i, nil
}

// IsValid 判断字符串是否为空
func IsValid(s string) error {
	if s == "" {
		return config.InvalidParameterErr
	}
	return nil
}

// StringToFloat string转为float
func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 32)
}
