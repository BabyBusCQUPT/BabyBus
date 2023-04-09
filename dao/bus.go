package dao

import (
	"BabyBus/config"
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

// InitGEO 初始化站点距离
func InitGEO(stationName string, longitude float64, latitude float64) error {
	err := RDB.GeoAdd("siteDistance", &redis.GeoLocation{
		Name:      stationName,
		Longitude: longitude,
		Latitude:  latitude,
	}).Err()
	return err
}

// UserSurroundings 用户周边站点
func UserSurroundings(userLongitude float64, userLatitude float64) ([]redis.GeoLocation, error) {
	res, err := RDB.GeoRadius("siteDistance", userLongitude, userLatitude, &redis.GeoRadiusQuery{
		Radius:      1000,
		Unit:        "m",
		WithCoord:   true,
		WithDist:    false,
		WithGeoHash: false,
		Count:       10,
		Sort:        "ASC",
	}).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func StationsScoreIncr(IncrNum float64, stationName string) error {
	err := RDB.ZIncrBy("hotStations", IncrNum, stationName).Err()
	return err
}

func GetHotStations() []string {
	hot := RDB.ZRevRangeByScore("hotStations", redis.ZRangeBy{
		Max: "+inf",
		Min: "-inf",
	}).Val()
	return hot
}

func CheckLimit(ip string) (int64, error) {
	count, err := RDB.Get(ip).Int64()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	return count, err
}

func IpIncrBy(ip string) error {
	_, err := RDB.IncrBy(ip, config.Incr).Result()
	return err
}

func TimeExpire(ip string) error {
	return RDB.Expire(ip, config.Expire).Err()
}
