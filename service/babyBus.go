package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func GetPersonalScore(bus *model.BabyBus) ([]model.BabyBus, error) {
	return dao.GetScore(bus)
}
