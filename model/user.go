package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OpenId     string
	SessionKey string
	Age        int
	Gender     byte
	Token      string
	Nickname   string
	Image      string
	MostUsed   int
	Friend     uint
}

type Position struct {
	Longitude float64
	Latitude  float64
}
