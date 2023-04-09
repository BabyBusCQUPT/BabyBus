package dao

import (
	"BabyBus/model"
	"github.com/go-redis/redis"
)

func GetBusScore(busId string) (sumScore float32, sumBaby int64, err error) {
	if err = DB.Model(&model.BabyBus{}).Where("busId = ?", busId).Pluck("SUM(sumScore) as sumScore", &sumScore).Error; err != nil {
		return -1, -1, err
	}
	if err = DB.Model(&model.BabyBus{}).Where("busId = ?", busId).Count(&sumBaby).Error; err != nil {
		return -1, -1, err
	}
	return sumScore, sumBaby, nil
}

func LikeSelect(keyWords string) (stations []model.Station, err error) {
	if err = DB.Model(&model.Station{}).Where("name LIKE " + "%" + keyWords).Find(&stations).Error; err != nil {
		return nil, err
	}
	return stations, nil
}

func GetStationDetails(stationName string) (station *model.Station, err error) {
	if err = DB.Model(&model.Station{}).Where("name = ?", stationName).Find(&station).Error; err != nil {
		return nil, err
	}
	return station, nil
}

func HotStations(stationName string) error {
	err := RDB.ZAdd("hotStations", redis.Z{Score: 0, Member: stationName}).Err()
	return err
}

func StationsScoreIncr(IncrNum float64, stationName string) error {
	err := RDB.ZIncrBy("hotStations", IncrNum, stationName).Err()
	return err
}
