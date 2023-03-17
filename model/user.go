package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OpenId        string
	SessionKey    string
	Age           int
	Gender        byte
	Token         string
	NikeName      string
	Image         string
	UserLongitude float32
	UserLatitude  float32
	MostUsed      int
}
