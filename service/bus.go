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

func GetStationDetails(stationName string) (*model.Station, error) {
	return dao.GetStationDetails(stationName)
}

func StationsScoreIncr(IncrNum float64, stationName string) error {
	return dao.StationsScoreIncr(IncrNum, stationName)
}

func GetHot() []string {
	return dao.GetHotStations()
}
