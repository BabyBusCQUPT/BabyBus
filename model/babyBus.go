package model

import "gorm.io/gorm"

type BabyBus struct {
	gorm.Model
	BabyId string
	BusId  int
	Score  float64
}
