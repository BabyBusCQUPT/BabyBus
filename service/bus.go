package service

import (
	"BabyBus/dao"
	"BabyBus/model"
)

func GetBusScore(busId string) (sumScore float32, sumBaby int64, err error) {
	return dao.GetBusScore(busId)
}

func SelectStations(keyWords string) ([]model.Station, error) {
	return dao.LikeSelect(keyWords)
}
